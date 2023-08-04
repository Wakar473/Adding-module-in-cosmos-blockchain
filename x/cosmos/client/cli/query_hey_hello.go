package cli

import (
	"strconv"

	"cosmos/x/cosmos/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdHeyHello() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hey-hello [to] [wakar]",
		Short: "Query hey-hello",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqTo := args[0]
			reqWakar := args[1]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryHeyHelloRequest{

				To:    reqTo,
				Wakar: reqWakar,
			}

			res, err := queryClient.HeyHello(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
