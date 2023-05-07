package simulation_test

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gridironOne/gstake-native/x/lsgridiron/simulation"
)

func TestParamChanges(t *testing.T) {
	s := rand.NewSource(1)
	r := rand.New(s)

	expected := []struct {
		composedKey string
		key         string
		simValue    string
		subspace    string
	}{
		{"lsgridiron/WhitelistedValidators", "WhitelistedValidators", "[]", "lsgridiron"},
		{"lsgridiron/LiquidBondDenom", "LiquidBondDenom", "\"bstake\"", "lsgridiron"},
		{"lsgridiron/UnstakeFeeRate", "UnstakeFeeRate", "\"0.010000000000000000\"", "lsgridiron"},
		{"lsgridiron/MinLiquidStakingAmount", "MinLiquidStakingAmount", "\"9727887\"", "lsgridiron"},
	}

	paramChanges := simulation.ParamChanges(r)
	require.Len(t, paramChanges, 4)

	for i, p := range paramChanges {
		require.Equal(t, expected[i].composedKey, p.ComposedKey())
		require.Equal(t, expected[i].key, p.Key())
		require.Equal(t, expected[i].simValue, p.SimValue()(r))
		require.Equal(t, expected[i].subspace, p.Subspace())
	}
}
