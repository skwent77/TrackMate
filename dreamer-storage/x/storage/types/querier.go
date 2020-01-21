package types

import (
	"strings"
	"time"
)

type QueryResAddrs []string

func (n QueryResAddrs) String() string {
	return strings.Join(n[:], "\n")
}

type QueryResAllData Data

type QueryResRangeData map[time.Time]Data
