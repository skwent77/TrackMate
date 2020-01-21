package storage

import (
	"github.com/dreamer-epitech/dreamer-storage/x/storage/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewMsgSetData = types.NewMsgSetData
	ModuleCdc     = types.ModuleCdc
	RegisterCodec = types.RegisterCodec
)

type (
	MsgSetData        = types.MsgSetData
	QueryResAddrs     = types.QueryResAddrs
	QueryResData      = types.QueryResAllData
	QueryResRangeData = types.QueryResRangeData
)
