package cli

import (
	"fmt"
	// "strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/simplyvc/interchainqueries/x/icq/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group icq queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdListPendingICQRequests())
	cmd.AddCommand(CmdShowPendingICQRequest())
	cmd.AddCommand(CmdListPendingICQRequestsResult())
	cmd.AddCommand(CmdShowPendingICQRequestResult())
	cmd.AddCommand(CmdListPendingICQRequestsTimeouts())
	cmd.AddCommand(CmdShowPendingICQRequestTimeouts())
	cmd.AddCommand(CmdListPendingICQRequestsPeriodic())
	cmd.AddCommand(CmdShowPendingICQRequestPeriodic())
	cmd.AddCommand(CmdListDataPointResult())
	cmd.AddCommand(CmdShowDataPointResult())
	// this line is used by starport scaffolding # 1

	return cmd
}
