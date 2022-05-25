package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/simplyvc/interchainqueries/x/icq/types"
)

func CmdListPendingICQRequestsTimeouts() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-interchainquery-timeouts",
		Short: "list the timeout counts and last timeout heights for all periodic queries.",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllICQTimeoutsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ICQTimeoutsAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowPendingICQRequestTimeouts() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-interchainquery-timeouts [id]",
		Short: "shows the timeouts of a periodic interchainquery",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetICQTimeoutsRequest{
				Id: id,
			}

			res, err := queryClient.ICQTimeouts(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}