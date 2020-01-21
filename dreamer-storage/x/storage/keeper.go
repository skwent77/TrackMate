package storage

import (
	"errors"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	storeKey sdk.StoreKey
	cdc      *codec.Codec
}

func NewKeeper(storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,
	}
}

func (k Keeper) SetData(ctx sdk.Context, address sdk.AccAddress, timestamp time.Time, data string) error {
	var value map[time.Time]string

	// load kvstore
	store := ctx.KVStore(k.storeKey)

	if store.Has(address) {
		tmp := store.Get(address)
		k.cdc.MustUnmarshalBinaryBare(tmp, &value)
	} else {
		value = make(map[time.Time]string)
	}

	// put key and value into map
	value[timestamp] = data

	// set the map to kvstore
	store.Set(address, k.cdc.MustMarshalBinaryBare(value))

	return nil
}

func (k Keeper) DeleteData(ctx sdk.Context, address sdk.AccAddress) error {
	// load kvstore
	store := ctx.KVStore(k.storeKey)

	if !store.Has(address) {
		return errors.New("no data corresponding to address")
	}

	// delete the data
	store.Delete(address)

	return nil
}

func (k Keeper) GetAllData(ctx sdk.Context, address sdk.AccAddress) map[time.Time]string {
	var value map[time.Time]string

	// load kvstore
	store := ctx.KVStore(k.storeKey)

	if !store.Has(address) {
		return nil
	}

	// get data(value) with address(key) from kvstore
	tmp := store.Get(address)
	k.cdc.MustUnmarshalBinaryBare(tmp, &value)

	return value
}

func (k Keeper) GetRangeData(ctx sdk.Context, address sdk.AccAddress, from, to time.Time) map[time.Time]string {
	var (
		rawData    map[time.Time]string
		resultData map[time.Time]string
	)

	// load kvstore
	store := ctx.KVStore(k.storeKey)

	if !store.Has(address) {
		return nil
	}

	// get data(value) with address(key) from kvstore
	tmp := store.Get(address)
	k.cdc.MustUnmarshalBinaryBare(tmp, &rawData)

	resultData = make(map[time.Time]string)

	for key, value := range rawData {
		if (from.Equal(key) || from.After(key)) && (to.Equal(key) || to.Before(key)) {
			resultData[key] = value
		}
	}

	if len(resultData) == 0 {
		return nil
	}

	return resultData
}

func (k Keeper) GetAddrs(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte{})
}
