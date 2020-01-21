package storage

import (
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryAddrs     = "addrs"
	QueryData      = "all_data"
	QueryRangeData = "range_data"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryAddrs:
			return queryAddrs(ctx, path[1:], req, keeper)
		case QueryData:
			return queryAllData(ctx, path[1:], req, keeper)
		case QueryRangeData:
			return queryRangeData(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown storage query endpoint")
		}
	}
}

func queryAddrs(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	var addrList QueryResAddrs

	iterator := keeper.GetAddrs(ctx)

	for ; iterator.Valid(); iterator.Next() {
		addrList = append(addrList, string(iterator.Key()))
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, addrList)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}

func queryAllData(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	addr, err := sdk.AccAddressFromBech32(path[0])
	if err != nil {
		return nil, sdk.ErrInvalidAddress(path[0])
	}

	value := keeper.GetAllData(ctx, addr)
	if value == nil {
		return nil, sdk.ErrUnknownRequest("no data corresponding to address")
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, value)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}

func queryRangeData(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	var (
		fromTime, toTime time.Time
	)

	addr, err := sdk.AccAddressFromBech32(path[0])
	if err != nil {
		return nil, sdk.ErrInvalidAddress(path[0])
	}

	fromTime, err = time.Parse(time.UnixDate, path[1])
	if err != nil {
		return nil, sdk.ErrInternal("wrong from time type")
	}

	toTime, err = time.Parse(time.UnixDate, path[2])
	if err != nil {
		return nil, sdk.ErrInternal("wrong to time type")
	}

	value := keeper.GetRangeData(ctx, addr, fromTime, toTime)
	if value == nil {
		return nil, sdk.ErrUnknownRequest("no data corresponding to address")
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, value)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}
