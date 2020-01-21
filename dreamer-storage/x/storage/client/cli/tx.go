package cli

import (
	"fmt"
	"time"

	"github.com/dreamer-epitech/dreamer-storage/x/storage/types"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	storageTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Storage transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	storageTxCmd.AddCommand(client.PostCommands()...)

	return storageTxCmd
}

func GetTxSetData(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-data [address] [timestamp] [data]",
		Short: "set the value associated with a timestamp",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			addr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				fmt.Printf("wrong address type\n")
				return err
			}

			timestamp, err := time.Parse(time.UnixDate, args[1])
			if err != nil {
				fmt.Printf("error happened while parsing timestamp\n")
				return err
			}

			data := args[2]

			msg := types.NewMsgSetData(addr, timestamp, data)
			err = msg.ValidateBasic()
			if err != nil {
				fmt.Printf("error validating msg\n")
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
