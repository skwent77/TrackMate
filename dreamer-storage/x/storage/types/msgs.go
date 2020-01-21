package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName

type MsgSetData struct {
	Address   sdk.AccAddress `json:"address"`
	Timestamp time.Time      `json:"timestamp"`
	Data      string         `json:"data"`
}

func NewMsgSetData(address sdk.AccAddress, timestamp time.Time, data string) MsgSetData {
	return MsgSetData{
		Address:   address,
		Timestamp: timestamp,
		Data:      data,
	}
}

func (msg MsgSetData) Route() string { return RouterKey }

func (msg MsgSetData) Type() string { return "set_data" }

func (msg MsgSetData) ValidateBasic() sdk.Error {
	if msg.Address.Empty() {
		return sdk.ErrInvalidAddress(msg.Address.String())
	}

	if msg.Timestamp.IsZero() {
		return sdk.ErrUnknownRequest("Invalid Timestamp")
	}

	return nil
}

func (msg MsgSetData) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgSetData) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Address}
}
