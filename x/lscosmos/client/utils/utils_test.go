package utils

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/gridironOne/gstake-native/x/lscosmos/types"
)

func TestNewMinDepositAndFeeChangeJSON(t *testing.T) {
	propJSON := NewMinDepositAndFeeChangeJSON(
		"title",
		"description",
		"5",
		"0.0",
		"0.0",
		"0.0",
		"0.1",
		"1000stake")

	require.Equal(t, "title", propJSON.Title)
	require.Equal(t, "description", propJSON.Description)
	require.Equal(t, "5", propJSON.MinDeposit)
	require.Equal(t, "0.0", propJSON.GstakeDepositFee)
	require.Equal(t, "0.0", propJSON.GstakeRestakeFee)
	require.Equal(t, "0.0", propJSON.GstakeUnstakeFee)
	require.Equal(t, "1000stake", propJSON.Deposit)
}

func TestNewGstakeFeeAddressChangeProposalJSON(t *testing.T) {
	propJSON := NewGstakeFeeAddressChangeProposalJSON(
		"title",
		"description",
		"did:fury:g1pss7nxeh3f9md2vuxku8q99femnwdjtcpe9ky9",
		"1000stake")

	require.Equal(t, "title", propJSON.Title)
	require.Equal(t, "description", propJSON.Description)
	require.Equal(t, "did:fury:g1pss7nxeh3f9md2vuxku8q99femnwdjtcpe9ky9", propJSON.GstakeFeeAddress)
	require.Equal(t, "1000stake", propJSON.Deposit)
}

func TestNewAllowListedValidatorSetChangeProposalJSON(t *testing.T) {
	propJSON := NewAllowListedValidatorSetChangeProposalJSON(
		"title",
		"description",
		"1000stake",
		types.AllowListedValidators{
			AllowListedValidators: []types.AllowListedValidator{{
				ValidatorAddress: "Valaddr",
				TargetWeight:     sdk.OneDec(),
			}}},
	)

	require.Equal(t, "title", propJSON.Title)
	require.Equal(t, "description", propJSON.Description)
	require.Equal(t, "Valaddr", propJSON.AllowListedValidators.AllowListedValidators[0].ValidatorAddress)
	require.Equal(t, sdk.OneDec(), propJSON.AllowListedValidators.AllowListedValidators[0].TargetWeight)
	require.Equal(t, "1000stake", propJSON.Deposit)

}
