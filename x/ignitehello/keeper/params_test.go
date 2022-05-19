package keeper_test

import (
	"testing"

	testkeeper "github.com/albacg5/ignite-hello/testutil/keeper"
	"github.com/albacg5/ignite-hello/x/ignitehello/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.IgnitehelloKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
