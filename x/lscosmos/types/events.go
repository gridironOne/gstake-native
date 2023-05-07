package types

// IBC events
const (
	EventTypePacket            = "ics27_packet"
	EventTypeTimeout           = "timeout"
	EventTypeLiquidStake       = "liquid-stake"
	EventTypeRedeem            = "redeem"
	EventTypeLiquidUnstake     = "liquid-unstake"
	EventTypeClaim             = "claim"
	EventTypeJumpStart         = "jump-start"
	EventTypeRecreateICA       = "recreate-ica"
	EventTypeChangeModuleState = "change-module-state"
	EventTypeReportSlashing    = "report-slashing"
	EventTypePerformSlashing   = "perform-slashing"
	// this line is used by starport scaffolding # ibc/packet/event

	AttributeKeyAckSuccess         = "success"
	AttributeKeyAck                = "acknowledgement"
	AttributeKeyAckError           = "error"
	AttributeAmount                = "amount"
	AttributeAmountReceived        = "received"
	AttributeUnstakeAmount         = "undelegation-amount"
	AttributeGstakeDepositFee      = "gstake-deposit-fee"
	AttributeGstakeRedeemFee       = "gstake-redeem-fee"
	AttributeGstakeUnstakeFee      = "gstake-unstake-fee"
	AttributeDelegatorAddress      = "address"
	AttributeRewarderAddress       = "rewarder-address"
	AttributeClaimedAmount         = "claimed-amount"
	AttributeGstakeAddress         = "gstake-address"
	AttributeFromAddress           = "from-address"
	AttributeRecreateDelegationICA = "recreate-delegation-ica"
	AttributeRecreateRewardsICA    = "recreate-rewards-ica"
	AttributeChangedModuleState    = "module-state"
	AttributeValidatorAddress      = "validator-address"
	AttributeExistingDelegation    = "existing-delegation"
	AttributeUpdatedDelegation     = "updated-delegation"
	AttributeSlashedAmount         = "slashed-amount"
	AttributeValueCategory         = ModuleName
)
