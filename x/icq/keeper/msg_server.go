package keeper

import (
	"context"
	"strconv"
	"strings"

	ics23 "github.com/confio/ics23/go"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	commitmenttypes "github.com/cosmos/ibc-go/v3/modules/core/23-commitment/types"

	"github.com/simplyvc/interchainqueries/x/icq/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// SubmitICQResult submits a result for a given ICQ instance request
func (k msgServer) SubmitICQResult(
	goCtx context.Context,
	msg *types.MsgSubmitICQResult,
) (*types.MsgSubmitICQResultResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)
	currHeight := uint64(ctx.BlockHeight())

	// Retrieve the given ICQ instance that we want to service to validate its existence
	icqInstance, found := k.GetPendingICQInstance(ctx, msg.QueryId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrICQNotFound, "(%d) not found", msg.QueryId)
	}

	// Retrieve the periodic ICQ for the full details of the query
	periodicICQ, found := k.GetPeriodicICQ(ctx, msg.PeriodicId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrPeriodicICQNotFound, "(%d) not found", msg.PeriodicId)
	}

	if icqInstance.PeriodicId != msg.PeriodicId {
		return nil, sdkerrors.Wrapf(types.ErrPeriodicIdNoMatch, "stored: (%d) doesn't match submitted (%d)",
			icqInstance.PeriodicId, msg.PeriodicId)
	}

	// Retrieve the consensus state to validate the proofs
	consensusState, found := k.clientKeeper.GetClientConsensusState(ctx, periodicICQ.ClientId, msg.Height)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrClientConsensusNotFound, "(%s) not found at height (%d)",
			periodicICQ.ClientId, msg.Height.RevisionHeight)
	}

	merkleProof, err := commitmenttypes.ConvertProofs(msg.Proof)
	if err != nil {
		return nil, err
	}

	// NOTE: merkleProof.Proofs[0] is the IavlSpec Proof
	// 		 merkleProof.Proofs[1] is the TendermintSpec Proof
	iavlRoot, err := merkleProof.Proofs[0].Calculate()
	if err != nil {
		return nil, err
	}

	// Verify if the IavlSpec Root exists in the TendermintSpec Proof.
	rootVerified := ics23.VerifyMembership(
		ics23.TendermintSpec,
		consensusState.GetRoot().GetHash(),
		merkleProof.Proofs[1],
		merkleProof.Proofs[1].GetExist().Key,
		iavlRoot,
	)
	if !rootVerified {
		return nil, sdkerrors.Wrapf(types.ErrIavlRootVerification, "root not verified")
	}

	// Verify that the value exists for the stored key using the verified iavlRoot
	kvVerified := ics23.VerifyMembership(
		ics23.IavlSpec,
		iavlRoot,
		merkleProof.Proofs[0],
		periodicICQ.QueryParameters,
		msg.Result,
	)
	if !kvVerified {
		return nil, sdkerrors.Wrapf(types.ErrKVVerification, "value: (%s) not verified for key (%s)", msg.Result,
			periodicICQ.QueryParameters)
	}

	// Set this as the last dataPoint and overwrite later based on conditions
	dataPointId := strconv.FormatUint(msg.PeriodicId, 10) + "/1"
	prevDataPointId := dataPointId

	// If the periodic result is not found, it means it's the first result
	periodicResult, found := k.GetPeriodicLastDataPointId(ctx, msg.PeriodicId)
	if found {

		// This should be found otherwise the initial periodic query doesn't exist
		prevDataPoint, _ := k.GetDataPoint(ctx, periodicResult)
		if prevDataPoint.TargetHeight >= msg.Height.RevisionHeight {
			return nil, sdkerrors.Wrapf(types.ErrDuplicateHeightSubmissions,
				"height: (%d) already has a value submitted", msg.Height)
		}
		prevDataPointId = prevDataPoint.Id

		// We create the new key to store the data point under
		splitResult := strings.Split(periodicResult, "/")
		prevResultId, err := strconv.ParseUint(splitResult[1], 10, 64)
		if err != nil {
			return nil, sdkerrors.Wrapf(types.ErrDataPointKeyBroken, "periodic id: (%s)", msg.PeriodicId)
		}

		// If the prevResultId is at the max amount of results for the periodic query we need to restart the loop
		if prevResultId != periodicICQ.MaxResults {
			dataPointId = strconv.FormatUint(msg.PeriodicId, 10) + "/" + strconv.FormatUint(prevResultId+1, 10)
		}
	}

	k.SetDataPoint(ctx, types.DataPoint{
		Id:              dataPointId,
		LocalHeight:     currHeight,
		TargetHeight:    msg.Height.RevisionHeight,
		Data:            msg.Result,
		PrevDataPointId: prevDataPointId,
	})

	k.SetPeriodicLastDataPointId(ctx, msg.PeriodicId, dataPointId)

	k.RemovePendingICQInstance(ctx, msg.QueryId)

	return &types.MsgSubmitICQResultResponse{}, nil
}
