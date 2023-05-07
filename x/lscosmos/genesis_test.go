package lscosmos_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gridironOne/gstake-native/app/helpers"
	"github.com/gridironOne/gstake-native/x/lscosmos"
	"github.com/gridironOne/gstake-native/x/lscosmos/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # genesis/test/state
	}

	_, gStakeApp, ctx := helpers.CreateTestApp(t)
	k := gStakeApp.LSCosmosKeeper
	lscosmos.InitGenesis(ctx, k, genesisState)
	got := lscosmos.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
