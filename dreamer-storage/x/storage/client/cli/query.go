package cli

import (
	"fmt"

	"github.com/dreamer-epitech/dreamer-storage/x/storage/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	storageQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the storage module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE: client.ValidateCmd,
	}

	storageQueryCmd.AddCommand(client.GetCommands(
		GetCmdAddrs(storeKey, cdc),
		GetCmdAllData(storeKey, cdc),
		GetCmdRangeData(storeKey, cdc),
	)...)

	return storageQueryCmd
}

func GetCmdAddrs(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "addrs",
		Short: "addrs",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/addrs", queryRoute), nil)
			if err != nil {
				fmt.Printf("could not get query names\n")
				return nil
			}

			var out types.QueryResAddrs
			cdc.MustUnmarshalJSON(res, &out)

			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdAllData(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "all-data [address]",
		Short: "get all data of 'address'",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			addr := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/all_data/%s", queryRoute, addr), nil)
			if err != nil {
				fmt.Printf("could not find data corresponding to %s\n", addr)
				return nil
			}

			fmt.Printf("%s\n", res)

			return nil
		},
	}
}

func GetCmdRangeData(QueryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "range-data [address] [from_timestamp] [to_timestamp]",
		Short: "get all data of 'address' between 'from' and 'to' unix-timestamps",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			addr := args[0]
			fromTimestamp := args[1]
			toTimestamp := args[2]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/range_data/%s/%s/%s", QueryRoute, addr, fromTimestamp, toTimestamp), nil)
			if err != nil {
				fmt.Printf("could not find data corresponding to %s within the timestamps\n", addr)
				return nil
			}

			fmt.Printf("%s\n", res)

			return nil
		},
	}
}
