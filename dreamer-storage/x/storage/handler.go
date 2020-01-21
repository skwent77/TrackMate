package storage

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgSetData:
			return handleMsgSetData(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized storage Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgSetData(ctx sdk.Context, keeper Keeper, msg MsgSetData) sdk.Result {
	keeper.SetData(ctx, msg.Address, msg.Timestamp, msg.Data)
	return sdk.Result{} // return
}
