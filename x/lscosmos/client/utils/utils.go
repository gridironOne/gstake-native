package utils

import (
	"os"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/gridironOne/gstake-native/v2/x/lscosmos/types"
)

// GstakeParams defines the fees and address for register host chain proposal's GstakeParams
type GstakeParams struct {
	GstakeDepositFee    string `json:"gstake_deposit_fee" yaml:"gstake_deposit_fee"`
	GstakeRestakeFee    string `json:"gstake_restake_fee" yaml:"gstake_restake_fee"`
	GstakeUnstakeFee    string `json:"gstake_unstake_fee" yaml:"gstake_unstake_fee"`
	GstakeRedemptionFee string `json:"gstake_redemption_fee" yaml:"gstake_redemption_fee"`
	GstakeFeeAddress    string `json:"gstake_fee_address" yaml:"gstake_fee_address"`
}

// MinDepositAndFeeChangeProposalJSON defines a MinDepositAndFeeChangeProposal JSON input to be parsed
// from a JSON file. Deposit is used by gov module to change status of proposal.
type MinDepositAndFeeChangeProposalJSON struct {
	Title               string `json:"title" yaml:"title"`
	Description         string `json:"description" yaml:"description"`
	MinDeposit          string `json:"min_deposit" yaml:"min_deposit"`
	GstakeDepositFee    string `json:"gstake_deposit_fee" yaml:"gstake_deposit_fee"`
	GstakeRestakeFee    string `json:"gstake_restake_fee" yaml:"gstake_restake_fee"`
	GstakeUnstakeFee    string `json:"gstake_unstake_fee" yaml:"gstake_unstake_fee"`
	GstakeRedemptionFee string `json:"gstake_redemption_fee" yaml:"gstake_redemption_fee"`
	Deposit             string `json:"deposit" yaml:"deposit"`
}

// NewMinDepositAndFeeChangeJSON returns MinDepositAndFeeChangeProposalJSON struct with input values
func NewMinDepositAndFeeChangeJSON(title, description, minDeposit, gstakeDepositFee, gstakeRestakeFee,
	gstakeUnstakeFee, gstakeRedemptionFee, deposit string) MinDepositAndFeeChangeProposalJSON {
	return MinDepositAndFeeChangeProposalJSON{
		Title:               title,
		Description:         description,
		MinDeposit:          minDeposit,
		GstakeDepositFee:    gstakeDepositFee,
		GstakeRestakeFee:    gstakeRestakeFee,
		GstakeUnstakeFee:    gstakeUnstakeFee,
		GstakeRedemptionFee: gstakeRedemptionFee,
		Deposit:             deposit,
	}

}

// ParseMinDepositAndFeeChangeProposalJSON reads and parses a MinDepositAndFeeChangeProposalJSON from
// file.
func ParseMinDepositAndFeeChangeProposalJSON(cdc *codec.LegacyAmino, proposalFile string) (MinDepositAndFeeChangeProposalJSON, error) {
	proposal := MinDepositAndFeeChangeProposalJSON{}

	contents, err := os.ReadFile(proposalFile)
	if err != nil {
		return proposal, err
	}
	if err := cdc.UnmarshalJSON(contents, &proposal); err != nil {
		return proposal, err
	}

	return proposal, nil
}

// GstakeFeeAddressChangeProposalJSON defines a GstakeFeeAddressChangeProposal JSON input to be parsed
// from a JSON file. Deposit is used by gov module to change status of proposal.
type GstakeFeeAddressChangeProposalJSON struct {
	Title            string `json:"title" yaml:"title"`
	Description      string `json:"description" yaml:"description"`
	GstakeFeeAddress string `json:"gstake_fee_address" yaml:"gstake_fee_address"`
	Deposit          string `json:"deposit" yaml:"deposit"`
}

// NewGstakeFeeAddressChangeProposalJSON returns GstakeFeeAddressChangeProposalJSON struct with input values
func NewGstakeFeeAddressChangeProposalJSON(title, description, gstakeFeeAddress, deposit string) GstakeFeeAddressChangeProposalJSON {
	return GstakeFeeAddressChangeProposalJSON{
		Title:            title,
		Description:      description,
		GstakeFeeAddress: gstakeFeeAddress,
		Deposit:          deposit,
	}

}

// ParseGstakeFeeAddressChangeProposalJSON reads and parses a GstakeFeeAddressChangeProposal  from
// file.
func ParseGstakeFeeAddressChangeProposalJSON(cdc *codec.LegacyAmino, proposalFile string) (GstakeFeeAddressChangeProposalJSON, error) {
	proposal := GstakeFeeAddressChangeProposalJSON{}

	contents, err := os.ReadFile(proposalFile)
	if err != nil {
		return proposal, err
	}
	if err := cdc.UnmarshalJSON(contents, &proposal); err != nil {
		return proposal, err
	}

	return proposal, nil
}

// JumpstartTxnJSON defines a JumpStart JSON input to be parsed
// from a JSON file.
type AllowListedValidatorSetChangeProposalJSON struct {
	Title                 string                      `json:"title" yaml:"title"`
	Description           string                      `json:"description" yaml:"description"`
	AllowListedValidators types.AllowListedValidators `json:"allow_listed_validators" yaml:"allow_listed_validators"`
	Deposit               string                      `json:"deposit" yaml:"deposit"`
}

// NewAllowListedValidatorSetChangeProposalJSON returns AllowListedValidatorSetChangeProposalJSON struct with input values
func NewAllowListedValidatorSetChangeProposalJSON(title, description, deposit string, allowListedValidators types.AllowListedValidators) AllowListedValidatorSetChangeProposalJSON {
	return AllowListedValidatorSetChangeProposalJSON{
		Title:                 title,
		Description:           description,
		AllowListedValidators: allowListedValidators,
		Deposit:               deposit,
	}

}

// ParseAllowListedValidatorSetChangeProposalJSON  reads and parses a AllowListedValidatorSetChangeProposalJSON  from
// file.
func ParseAllowListedValidatorSetChangeProposalJSON(cdc *codec.LegacyAmino, proposalFile string) (AllowListedValidatorSetChangeProposalJSON, error) {
	proposal := AllowListedValidatorSetChangeProposalJSON{}

	contents, err := os.ReadFile(proposalFile)
	if err != nil {
		return proposal, err
	}
	if err := cdc.UnmarshalJSON(contents, &proposal); err != nil {
		return proposal, err
	}

	return proposal, nil
}

// JumpstartTxnJSON defines a Jump start JSON input to be parsed
// from a JSON file.
type JumpstartTxnJSON struct {
	ChainID               string                      `json:"chain_id" yaml:"chain_id"`
	ConnectionID          string                      `json:"connection_id" yaml:"connection_id"`
	TransferChannel       string                      `json:"transfer_channel" yaml:"transfer_channel"`
	TransferPort          string                      `json:"transfer_port" yaml:"transfer_port"`
	BaseDenom             string                      `json:"base_denom" yaml:"base_denom"`
	MintDenom             string                      `json:"mint_denom" yaml:"mint_denom"`
	MinDeposit            string                      `json:"min_deposit" yaml:"min_deposit"`
	AllowListedValidators types.AllowListedValidators `json:"allow_listed_validators" yaml:"allow_listed_validators"`
	GstakeParams          GstakeParams                `json:"gstake_params" yaml:"gstake_params"`
	HostAccounts          types.HostAccounts          `json:"host_accounts" yaml:"host_accounts"`
}

// ParseJumpstartTxnJSON  reads and parses a JumpstartTxnJSON  from
// file.
func ParseJumpstartTxnJSON(cdc *codec.LegacyAmino, file string) (JumpstartTxnJSON, error) {
	jsonTxn := JumpstartTxnJSON{}

	contents, err := os.ReadFile(file)
	if err != nil {
		return jsonTxn, err
	}
	if err := cdc.UnmarshalJSON(contents, &jsonTxn); err != nil {
		return jsonTxn, err
	}

	return jsonTxn, nil
}
