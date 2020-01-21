package storage

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	data map[time.Time]string `json:"data"`
}

func NewGenesisState(whoIsRecords map[time.Time]string) GenesisState {
	return GenesisState{data: nil}
}

func ValidateGenesis(value GenesisState) error {
	// pass all

	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		data: map[time.Time]string{},
	}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, value GenesisState) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	return GenesisState{data: nil}
}
