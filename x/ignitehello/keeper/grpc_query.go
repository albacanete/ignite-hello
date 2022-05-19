package keeper

import (
	"github.com/albacg5/ignite-hello/x/ignitehello/types"
)

var _ types.QueryServer = Keeper{}
