package types

import (
	"fmt"
	"strings"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	ProposalTypeMinDepositAndFeeChange    = "MinDepositAndFeeChange"
	ProposalGstakeFeeAddressChange        = "GstakeFeeAddressChange"
	ProposalAllowListedValidatorSetChange = "AllowListedValidatorSetChange"
)

var (
	_ govtypes.Content = &MinDepositAndFeeChangeProposal{}
	_ govtypes.Content = &GstakeFeeAddressChangeProposal{}
	_ govtypes.Content = &AllowListedValidatorSetChangeProposal{}
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeMinDepositAndFeeChange)
	govtypes.RegisterProposalType(ProposalGstakeFeeAddressChange)
	govtypes.RegisterProposalType(ProposalAllowListedValidatorSetChange)
}

// NewHostChainParams returns HostChainParams with the input provided
func NewHostChainParams(chainID, connectionID, channel, port, baseDenom, mintDenom, gstakefeeAddress string, minDeposit math.Int, gstakeDepositFee, gstakeRestakeFee, gstakeUnstakeFee, gstakeRedemptionFee sdktypes.Dec) HostChainParams {
	return HostChainParams{
		ChainID:         chainID,
		ConnectionID:    connectionID,
		TransferChannel: channel,
		TransferPort:    port,
		BaseDenom:       baseDenom,
		MintDenom:       mintDenom,
		MinDeposit:      minDeposit,
		GstakeParams: GstakeParams{
			GstakeDepositFee:    gstakeDepositFee,
			GstakeRestakeFee:    gstakeRestakeFee,
			GstakeUnstakeFee:    gstakeUnstakeFee,
			GstakeRedemptionFee: gstakeRedemptionFee,
			GstakeFeeAddress:    gstakefeeAddress,
		},
	}
}

// IsEmpty Checks if HostChainParams were initialised
func (c *HostChainParams) IsEmpty() bool {
	if c.TransferChannel == "" ||
		c.TransferPort == "" ||
		c.ConnectionID == "" ||
		c.ChainID == "" ||
		c.BaseDenom == "" ||
		c.MintDenom == "" ||
		c.GstakeParams.GstakeFeeAddress == "" {
		return true
	}
	// can add more, but this should be good enough

	return false
}

// NewMinDepositAndFeeChangeProposal creates a protocol fee and min deposit change proposal.
func NewMinDepositAndFeeChangeProposal(title, description string, minDeposit math.Int, gstakeDepositFee,
	gstakeRestakeFee, gstakeUnstakeFee, gstakeRedemptionFee sdktypes.Dec) *MinDepositAndFeeChangeProposal {

	return &MinDepositAndFeeChangeProposal{
		Title:               title,
		Description:         description,
		MinDeposit:          minDeposit,
		GstakeDepositFee:    gstakeDepositFee,
		GstakeRestakeFee:    gstakeRestakeFee,
		GstakeUnstakeFee:    gstakeUnstakeFee,
		GstakeRedemptionFee: gstakeRedemptionFee,
	}
}

// GetTitle returns the title of the min-deposit and fee change proposal.
func (m *MinDepositAndFeeChangeProposal) GetTitle() string {
	return m.Title
}

// GetDescription returns the description of the min-deposit and fee change proposal.
func (m *MinDepositAndFeeChangeProposal) GetDescription() string {
	return m.Description
}

// ProposalRoute returns the proposal-route of the min-deposit and fee change proposal.
func (m *MinDepositAndFeeChangeProposal) ProposalRoute() string {
	return RouterKey
}

// ProposalType returns the proposal-type of the min-deposit and fee change proposal.
func (m *MinDepositAndFeeChangeProposal) ProposalType() string {
	return ProposalTypeMinDepositAndFeeChange
}

// ValidateBasic runs basic stateless validity checks
func (m *MinDepositAndFeeChangeProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(m)
	if err != nil {
		return err
	}

	if m.GstakeDepositFee.IsNegative() || m.GstakeDepositFee.GTE(MaxGstakeDepositFee) {
		return errorsmod.Wrapf(ErrInvalidFee, "gstake deposit fee must be between %s and %s", sdktypes.ZeroDec(), MaxGstakeDepositFee)
	}

	if m.GstakeRestakeFee.IsNegative() || m.GstakeRestakeFee.GTE(MaxGstakeRestakeFee) {
		return errorsmod.Wrapf(ErrInvalidFee, "gstake restake fee must be between %s and %s", sdktypes.ZeroDec(), MaxGstakeRestakeFee)
	}

	if m.GstakeUnstakeFee.IsNegative() || m.GstakeUnstakeFee.GTE(MaxGstakeUnstakeFee) {
		return errorsmod.Wrapf(ErrInvalidFee, "gstake unstake fee must be between %s and %s", sdktypes.ZeroDec(), MaxGstakeUnstakeFee)
	}

	if m.GstakeRedemptionFee.IsNegative() || m.GstakeRedemptionFee.GTE(MaxGstakeRedemptionFee) {
		return errorsmod.Wrapf(ErrInvalidFee, "gstake redemption fee must be between %s and %s", sdktypes.ZeroDec(), MaxGstakeRedemptionFee)
	}

	if m.MinDeposit.LTE(sdktypes.ZeroInt()) {
		return errorsmod.Wrapf(ErrInvalidDeposit, "min deposit must be positive")
	}

	return nil
}

// String returns the string of proposal details
func (m *MinDepositAndFeeChangeProposal) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`MinDepositAndFeeChange:
Title:                 %s
Description:           %s
MinDeposit:             %s
GstakeDepositFee:	   %s
GstakeRestakeFee: 	   %s
GstakeUnstakeFee: 	   %s
GstakeRedemptionFee:   %s

`,
		m.Title,
		m.Description,
		m.MinDeposit,
		m.GstakeDepositFee,
		m.GstakeRestakeFee,
		m.GstakeUnstakeFee,
		m.GstakeRedemptionFee),
	)
	return b.String()
}

// NewGstakeFeeAddressChangeProposal creates a gstake fee  address change proposal.
func NewGstakeFeeAddressChangeProposal(title, description,
	gstakeFeeAddress string) *GstakeFeeAddressChangeProposal {
	return &GstakeFeeAddressChangeProposal{
		Title:            title,
		Description:      description,
		GstakeFeeAddress: gstakeFeeAddress,
	}
}

// GetTitle returns the title of fee collector gstake fee address change proposal.
func (m *GstakeFeeAddressChangeProposal) GetTitle() string {
	return m.Title
}

// GetDescription returns the description of the gstake fee address proposal.
func (m *GstakeFeeAddressChangeProposal) GetDescription() string {
	return m.Description
}

// ProposalRoute returns the proposal-route of gstake fee address proposal.
func (m *GstakeFeeAddressChangeProposal) ProposalRoute() string {
	return RouterKey
}

// ProposalType returns the proposal-type of gstake fee address change proposal.
func (m *GstakeFeeAddressChangeProposal) ProposalType() string {
	return ProposalGstakeFeeAddressChange
}

// ValidateBasic runs basic stateless validity checks
func (m *GstakeFeeAddressChangeProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(m)
	if err != nil {
		return err
	}

	_, err = sdktypes.AccAddressFromBech32(m.GstakeFeeAddress)
	if err != nil {
		return err
	}

	return nil
}

// String returns the string of proposal details
func (m *GstakeFeeAddressChangeProposal) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`GstakeFeeAddressChange:
Title:                 %s
Description:           %s
GstakeFeeAddress: 	   %s

`,
		m.Title,
		m.Description,
		m.GstakeFeeAddress,
	),
	)
	return b.String()
}

// NewAllowListedValidatorSetChangeProposal creates a allowListed validator set change proposal.
func NewAllowListedValidatorSetChangeProposal(title, description string, allowListedValidators AllowListedValidators) *AllowListedValidatorSetChangeProposal {
	return &AllowListedValidatorSetChangeProposal{
		Title:                 title,
		Description:           description,
		AllowListedValidators: allowListedValidators,
	}
}

// GetTitle returns the title of allowListed validator set change proposal.
func (m *AllowListedValidatorSetChangeProposal) GetTitle() string {
	return m.Title
}

// GetDescription returns the description of allowListed validator set change proposal.
func (m *AllowListedValidatorSetChangeProposal) GetDescription() string {
	return m.Description
}

// ProposalRoute returns the proposal-route of allowListed validator set change proposal.
func (m *AllowListedValidatorSetChangeProposal) ProposalRoute() string {
	return RouterKey
}

// ProposalType returns the proposal-type of allowListed validator set change proposal.
func (m *AllowListedValidatorSetChangeProposal) ProposalType() string {
	return ProposalAllowListedValidatorSetChange
}

// ValidateBasic runs basic stateless validity checks
func (m *AllowListedValidatorSetChangeProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(m)
	if err != nil {
		return err
	}

	if !m.AllowListedValidators.Valid() {
		return errorsmod.Wrapf(ErrInValidAllowListedValidators, "allow listed validators is not valid")
	}

	return nil
}

// String returns the string of proposal details
func (m *AllowListedValidatorSetChangeProposal) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`AllowListedValidatorSetChange:
Title:                 %s
Description:           %s
AllowListedValidators: 	   %s

`,
		m.Title,
		m.Description,
		m.AllowListedValidators,
	),
	)
	return b.String()
}
