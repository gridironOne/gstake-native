package types_test

import (
	"fmt"
	"strings"
	"testing"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/stretchr/testify/require"

	"github.com/gridironOne/gstake-native/v2/app"
	"github.com/gridironOne/gstake-native/v2/x/lscosmos/types"
)

func init() {
	app.SetAddressPrefixes()
}

func TestNewMinDepositAndFeeChangeProposal(t *testing.T) {
	testCases := []struct {
		testName, expectedString string
		proposal                 types.MinDepositAndFeeChangeProposal
		expectedError            error
	}{
		{
			testName: "correct proposal content",
			proposal: *types.NewMinDepositAndFeeChangeProposal(
				"title",
				"description",
				sdk.OneInt().MulRaw(5),
				sdk.ZeroDec(),
				sdk.ZeroDec(),
				sdk.ZeroDec(),
				sdk.ZeroDec(),
			),
			expectedError:  nil,
			expectedString: "MinDepositAndFeeChange:\nTitle:                 title\nDescription:           description\nMinDeposit:             5\nGstakeDepositFee:\t   0.000000000000000000\nGstakeRestakeFee: \t   0.000000000000000000\nGstakeUnstakeFee: \t   0.000000000000000000\nGstakeRedemptionFee:   0.000000000000000000\n\n",
		},
		{
			testName: "invalid title length",
			proposal: *types.NewMinDepositAndFeeChangeProposal(
				"",
				"description",
				sdk.OneInt().MulRaw(5),
				sdk.ZeroDec(),
				sdk.ZeroDec(),
				sdk.ZeroDec(),
				sdk.ZeroDec(),
			),
			expectedError: errorsmod.Wrap(govtypes.ErrInvalidProposalContent, "proposal title cannot be blank"),
		},
		{
			testName: "invalid title length",
			proposal: *types.NewMinDepositAndFeeChangeProposal(
				strings.Repeat("-", govv1beta1.MaxTitleLength+1),
				"description",
				sdk.OneInt().MulRaw(5),
				sdk.ZeroDec(),
				sdk.ZeroDec(),
				sdk.ZeroDec(),
				sdk.ZeroDec(),
			),
			expectedError: errorsmod.Wrapf(govtypes.ErrInvalidProposalContent, "proposal title is longer than max length of %d", govv1beta1.MaxTitleLength),
		},
		{
			testName: "invalid description length",
			proposal: *types.NewMinDepositAndFeeChangeProposal(
				"title",
				"",
				sdk.OneInt().MulRaw(5),
				sdk.ZeroDec(),
				sdk.ZeroDec(),
				sdk.ZeroDec(),
				sdk.ZeroDec(),
			),
			expectedError: errorsmod.Wrap(govtypes.ErrInvalidProposalContent, "proposal description cannot be blank"),
		},
		{
			testName: "invalid description length",
			proposal: *types.NewMinDepositAndFeeChangeProposal(
				"title",
				strings.Repeat("-", govv1beta1.MaxDescriptionLength+1),
				sdk.OneInt().MulRaw(5),
				sdk.ZeroDec(),
				sdk.ZeroDec(),
				sdk.ZeroDec(),
				sdk.ZeroDec(),
			),
			expectedError: errorsmod.Wrapf(govtypes.ErrInvalidProposalContent, "proposal description is longer than max length of %d", govv1beta1.MaxDescriptionLength),
		},
		{
			testName: "incorrect gstake deposit fee",
			proposal: *types.NewMinDepositAndFeeChangeProposal(
				"title",
				"description",
				sdk.OneInt().MulRaw(5),
				sdk.NewDec(10),
				sdk.ZeroDec(),
				sdk.ZeroDec(),
				sdk.ZeroDec(),
			),
			expectedError: errorsmod.Wrapf(types.ErrInvalidFee, "gstake deposit fee must be between %s and %s", sdk.ZeroDec(), types.MaxGstakeDepositFee),
		},
		{
			testName: "incorrect gstake restake fee",
			proposal: *types.NewMinDepositAndFeeChangeProposal(
				"title",
				"description",
				sdk.OneInt().MulRaw(5),
				sdk.ZeroDec(),
				sdk.NewDec(10),
				sdk.ZeroDec(),
				sdk.ZeroDec(),
			),
			expectedError: errorsmod.Wrapf(types.ErrInvalidFee, "gstake restake fee must be between %s and %s", sdk.ZeroDec(), types.MaxGstakeRestakeFee),
		},
		{
			testName: "incorrect gstake unstake fee",
			proposal: *types.NewMinDepositAndFeeChangeProposal(
				"title",
				"description",
				sdk.OneInt().MulRaw(5),
				sdk.ZeroDec(),
				sdk.ZeroDec(),
				sdk.NewDec(10),
				sdk.ZeroDec(),
			),
			expectedError: errorsmod.Wrapf(types.ErrInvalidFee, "gstake unstake fee must be between %s and %s", sdk.ZeroDec(), types.MaxGstakeUnstakeFee),
		},
		{
			testName: "incorrect gstake unstake fee",
			proposal: *types.NewMinDepositAndFeeChangeProposal(
				"title",
				"description",
				sdk.OneInt().MulRaw(5),
				sdk.ZeroDec(),
				sdk.ZeroDec(),
				sdk.ZeroDec(),
				sdk.NewDec(10),
			),
			expectedError: errorsmod.Wrapf(types.ErrInvalidFee, "gstake redemption fee must be between %s and %s", sdk.ZeroDec(), types.MaxGstakeRedemptionFee),
		},
		{
			testName: "incorrect deposit",
			proposal: *types.NewMinDepositAndFeeChangeProposal(
				"title",
				"description",
				sdk.OneInt().MulRaw(-1),
				sdk.ZeroDec(),
				sdk.ZeroDec(),
				sdk.ZeroDec(),
				sdk.ZeroDec(),
			),
			expectedError: errorsmod.Wrapf(types.ErrInvalidDeposit, "min deposit must be positive"),
		},
	}

	for _, tc := range testCases {
		require.Equal(t, types.RouterKey, tc.proposal.ProposalRoute())
		require.Equal(t, types.ProposalTypeMinDepositAndFeeChange, tc.proposal.ProposalType())

		if tc.expectedError != nil {
			require.Equal(t, tc.expectedError.Error(), tc.proposal.ValidateBasic().Error())
		}
		if tc.expectedError == nil {
			require.Equal(t, "title", tc.proposal.GetTitle())
			require.Equal(t, "description", tc.proposal.GetDescription())
			require.Equal(t, tc.expectedString, tc.proposal.String())
		}
	}
}

func TestNewGstakeFeeAddressChangeProposal(t *testing.T) {
	testCases := []struct {
		testName, expectedString string
		proposal                 types.GstakeFeeAddressChangeProposal
		expectedError            error
	}{
		{
			testName: "correct proposal content",
			proposal: *types.NewGstakeFeeAddressChangeProposal(
				"title",
				"description",
				"gridiron1pss7nxeh3f9md2vuxku8q99femnwdjtcpe9ky9",
			),
			expectedError:  nil,
			expectedString: "GstakeFeeAddressChange:\nTitle:                 title\nDescription:           description\nGstakeFeeAddress: \t   gridiron1pss7nxeh3f9md2vuxku8q99femnwdjtcpe9ky9\n\n",
		},
		{
			testName: "invalid title length",
			proposal: *types.NewGstakeFeeAddressChangeProposal(
				"",
				"description",
				"gridiron1pss7nxeh3f9md2vuxku8q99femnwdjtcpe9ky9",
			),
			expectedError: errorsmod.Wrap(govtypes.ErrInvalidProposalContent, "proposal title cannot be blank"),
		},
		{
			testName: "invalid title length",
			proposal: *types.NewGstakeFeeAddressChangeProposal(
				strings.Repeat("-", govv1beta1.MaxTitleLength+1),
				"description",
				"gridiron1pss7nxeh3f9md2vuxku8q99femnwdjtcpe9ky9",
			),
			expectedError: errorsmod.Wrapf(govtypes.ErrInvalidProposalContent, "proposal title is longer than max length of %d", govv1beta1.MaxTitleLength),
		},
		{
			testName: "invalid description length",
			proposal: *types.NewGstakeFeeAddressChangeProposal(
				"title",
				"",
				"gridiron1pss7nxeh3f9md2vuxku8q99femnwdjtcpe9ky9",
			),
			expectedError: errorsmod.Wrap(govtypes.ErrInvalidProposalContent, "proposal description cannot be blank"),
		},
		{
			testName: "invalid description length",
			proposal: *types.NewGstakeFeeAddressChangeProposal(
				"title",
				strings.Repeat("-", govv1beta1.MaxDescriptionLength+1),
				"gridiron1pss7nxeh3f9md2vuxku8q99femnwdjtcpe9ky9",
			),
			expectedError: errorsmod.Wrapf(govtypes.ErrInvalidProposalContent, "proposal description is longer than max length of %d", govv1beta1.MaxDescriptionLength),
		},
		{
			testName: "invalid gstake fee address length",
			proposal: *types.NewGstakeFeeAddressChangeProposal(
				"title",
				"description",
				"cosmos1hcqg5wj9t42zawqkqucs7la85ffyv08lum327c",
			),
			expectedError: fmt.Errorf("invalid Bech32 prefix; expected gridiron, got cosmos"),
		},
	}
	for _, tc := range testCases {
		require.Equal(t, types.RouterKey, tc.proposal.ProposalRoute())
		require.Equal(t, types.ProposalGstakeFeeAddressChange, tc.proposal.ProposalType())

		if tc.expectedError != nil {
			require.Equal(t, tc.expectedError.Error(), tc.proposal.ValidateBasic().Error())
		}
		if tc.expectedError == nil {
			require.Equal(t, "title", tc.proposal.GetTitle())
			require.Equal(t, "description", tc.proposal.GetDescription())
			require.Equal(t, tc.expectedString, tc.proposal.String())
		}
	}
}

func TestNewAllowListedValidatorSetChangeProposal(t *testing.T) {
	testCases := []struct {
		testName, expectedString string
		proposal                 types.AllowListedValidatorSetChangeProposal
		expectedError            error
	}{
		{
			testName: "correct proposal content",
			proposal: *types.NewAllowListedValidatorSetChangeProposal(
				"title",
				"description",
				types.AllowListedValidators{AllowListedValidators: []types.AllowListedValidator{{ValidatorAddress: "cosmosvaloper1hcqg5wj9t42zawqkqucs7la85ffyv08le09ljt", TargetWeight: sdk.OneDec()}}},
			),
			expectedError:  nil,
			expectedString: "AllowListedValidatorSetChange:\nTitle:                 title\nDescription:           description\nAllowListedValidators: \t   {[{cosmosvaloper1hcqg5wj9t42zawqkqucs7la85ffyv08le09ljt 1.000000000000000000}]}\n\n",
		},
		{
			testName: "invalid title length",
			proposal: *types.NewAllowListedValidatorSetChangeProposal(
				"",
				"description",
				types.AllowListedValidators{AllowListedValidators: []types.AllowListedValidator{{ValidatorAddress: "cosmosvaloper1hcqg5wj9t42zawqkqucs7la85ffyv08le09ljt", TargetWeight: sdk.OneDec()}}},
			),
			expectedError: errorsmod.Wrap(govtypes.ErrInvalidProposalContent, "proposal title cannot be blank"),
		},
		{
			testName: "invalid title length",
			proposal: *types.NewAllowListedValidatorSetChangeProposal(
				strings.Repeat("-", govv1beta1.MaxTitleLength+1),
				"description",
				types.AllowListedValidators{AllowListedValidators: []types.AllowListedValidator{{ValidatorAddress: "cosmosvaloper1hcqg5wj9t42zawqkqucs7la85ffyv08le09ljt", TargetWeight: sdk.OneDec()}}},
			),
			expectedError: errorsmod.Wrapf(govtypes.ErrInvalidProposalContent, "proposal title is longer than max length of %d", govv1beta1.MaxTitleLength),
		},
		{
			testName: "invalid description length",
			proposal: *types.NewAllowListedValidatorSetChangeProposal(
				"title",
				"",
				types.AllowListedValidators{AllowListedValidators: []types.AllowListedValidator{{ValidatorAddress: "cosmosvaloper1hcqg5wj9t42zawqkqucs7la85ffyv08le09ljt", TargetWeight: sdk.OneDec()}}},
			),
			expectedError: errorsmod.Wrap(govtypes.ErrInvalidProposalContent, "proposal description cannot be blank"),
		},
		{
			testName: "invalid description length",
			proposal: *types.NewAllowListedValidatorSetChangeProposal(
				"title",
				strings.Repeat("-", govv1beta1.MaxDescriptionLength+1),
				types.AllowListedValidators{AllowListedValidators: []types.AllowListedValidator{{ValidatorAddress: "cosmosvaloper1hcqg5wj9t42zawqkqucs7la85ffyv08le09ljt", TargetWeight: sdk.OneDec()}}},
			),
			expectedError: errorsmod.Wrapf(govtypes.ErrInvalidProposalContent, "proposal description is longer than max length of %d", govv1beta1.MaxDescriptionLength),
		},
		{
			testName: "incorrect allow listed validators",
			proposal: *types.NewAllowListedValidatorSetChangeProposal(
				"title",
				"description",
				types.AllowListedValidators{AllowListedValidators: []types.AllowListedValidator{{ValidatorAddress: "cosmosvaloper1hcqg5wj9t42zawqkqucs7la85ffyv08le09ljt", TargetWeight: sdk.ZeroDec()}}},
			),
			expectedError: errorsmod.Wrapf(types.ErrInValidAllowListedValidators, "allow listed validators is not valid"),
		},
	}

	for _, tc := range testCases {
		require.Equal(t, types.RouterKey, tc.proposal.ProposalRoute())
		require.Equal(t, types.ProposalAllowListedValidatorSetChange, tc.proposal.ProposalType())

		if tc.expectedError != nil {
			require.Equal(t, tc.expectedError.Error(), tc.proposal.ValidateBasic().Error())
		}
		if tc.expectedError == nil {
			require.Equal(t, "title", tc.proposal.GetTitle())
			require.Equal(t, "description", tc.proposal.GetDescription())
			require.Equal(t, tc.expectedString, tc.proposal.String())
		}
	}
}
