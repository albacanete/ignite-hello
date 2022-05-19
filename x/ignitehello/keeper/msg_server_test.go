package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/albacg5/ignite-hello/testutil/keeper"
	"github.com/albacg5/ignite-hello/x/ignitehello/keeper"
	"github.com/albacg5/ignite-hello/x/ignitehello/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.IgnitehelloKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
