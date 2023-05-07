package utils

import (
	"os"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/gridironOne/gstake-native/v2/x/lscosmos/types"
)

// PstakeParams defines the fees and address for register host chain proposal's PstakeParams
type PstakeParams struct {
	PstakeDepositFee    string `json:"gstake_deposit_fee" yaml:"gstake_deposit_fee"`
	PstakeRestakeFee    string `json:"gstake_restake_fee" yaml:"gstake_restake_fee"`
	PstakeUnstakeFee    string `json:"gstake_unstake_fee" yaml:"gstake_unstake_fee"`
	PstakeRedemptionFee string `json:"gstake_redemption_fee" yaml:"gstake_redemption_fee"`
	PstakeFeeAddress    string `json:"gstake_fee_address" yaml:"gstake_fee_address"`
}

// MinDepositAndFeeChangeProposalJSON defines a MinDepositAndFeeChangeProposal JSON input to be parsed
// from a JSON file. Deposit is used by gov module to change status of proposal.
type MinDepositAndFeeChangeProposalJSON struct {
	Title               string `json:"title" yaml:"title"`
	Description         string `json:"description" yaml:"description"`
	MinDeposit          string `json:"min_deposit" yaml:"min_deposit"`
	PstakeDepositFee    string `json:"gstake_deposit_fee" yaml:"gstake_deposit_fee"`
	PstakeRestakeFee    string `json:"gstake_restake_fee" yaml:"gstake_restake_fee"`
	PstakeUnstakeFee    string `json:"gstake_unstake_fee" yaml:"gstake_unstake_fee"`
	PstakeRedemptionFee string `json:"gstake_redemption_fee" yaml:"gstake_redemption_fee"`
	Deposit             string `json:"deposit" yaml:"deposit"`
}

// NewMinDepositAndFeeChangeJSON returns MinDepositAndFeeChangeProposalJSON struct with input values
func NewMinDepositAndFeeChangeJSON(title, description, minDeposit, gstakeDepositFee, gstakeRestakeFee,
	gstakeUnstakeFee, gstakeRedemptionFee, deposit string) MinDepositAndFeeChangeProposalJSON {
	return MinDepositAndFeeChangeProposalJSON{
		Title:               title,
		Description:         description,
		MinDeposit:          minDeposit,
		PstakeDepositFee:    gstakeDepositFee,
		PstakeRestakeFee:    gstakeRestakeFee,
		PstakeUnstakeFee:    gstakeUnstakeFee,
		PstakeRedemptionFee: gstakeRedemptionFee,
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

// PstakeFeeAddressChangeProposalJSON defines a PstakeFeeAddressChangeProposal JSON input to be parsed
// from a JSON file. Deposit is used by gov module to change status of proposal.
type PstakeFeeAddressChangeProposalJSON struct {
	Title            string `json:"title" yaml:"title"`
	Description      string `json:"description" yaml:"description"`
	PstakeFeeAddress string `json:"gstake_fee_address" yaml:"gstake_fee_address"`
	Deposit          string `json:"deposit" yaml:"deposit"`
}

// NewPstakeFeeAddressChangeProposalJSON returns PstakeFeeAddressChangeProposalJSON struct with input values
func NewPstakeFeeAddressChangeProposalJSON(title, description, gstakeFeeAddress, deposit string) PstakeFeeAddressChangeProposalJSON {
	return PstakeFeeAddressChangeProposalJSON{
		Title:            title,
		Description:      description,
		PstakeFeeAddress: gstakeFeeAddress,
		Deposit:          deposit,
	}

}

// ParsePstakeFeeAddressChangeProposalJSON reads and parses a PstakeFeeAddressChangeProposal  from
// file.
func ParsePstakeFeeAddressChangeProposalJSON(cdc *codec.LegacyAmino, proposalFile string) (PstakeFeeAddressChangeProposalJSON, error) {
	proposal := PstakeFeeAddressChangeProposalJSON{}

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
	PstakeParams          PstakeParams                `json:"gstake_params" yaml:"gstake_params"`
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
