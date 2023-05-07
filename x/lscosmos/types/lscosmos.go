package types

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	icatypes "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/types"
)

type (
	HostAccountDelegations []HostAccountDelegation
	AllowListedVals        []AllowListedValidator
)

var (
	_ sort.Interface = &HostAccountDelegations{}
	_ sort.Interface = &AllowListedVals{}
)

// Valid performs validity checks on AllowListedValidators and returns bool
func (av *AllowListedValidators) Valid() bool {
	if av == nil {
		return false
	}
	if av.AllowListedValidators == nil || len(av.AllowListedValidators) == 0 {
		return false
	}

	noDuplicate := make(map[string]bool)
	sum := sdk.ZeroDec()
	for _, v := range av.AllowListedValidators {
		if _, err := ValAddressFromBech32(v.ValidatorAddress, CosmosValOperPrefix); err != nil {
			return false
		}
		_, ok := noDuplicate[v.ValidatorAddress]
		if ok {
			return false
		}
		noDuplicate[v.ValidatorAddress] = true

		if v.TargetWeight.IsNegative() {
			return false
		}

		sum = sum.Add(v.TargetWeight)
	}
	return sum.Equal(sdk.OneDec())
}

// GetAddressMap returns a map of AllowListedValidators address as key and weights as values
func GetAddressMap(validators AllowListedValidators) map[string]sdk.Dec {
	addressMap := map[string]sdk.Dec{}

	for _, val := range validators.AllowListedValidators {
		addressMap[val.ValidatorAddress] = val.TargetWeight
	}
	return addressMap
}

func (av AllowListedVals) Len() int {
	return len(av)
}

func (av AllowListedVals) Less(i, j int) bool {
	return av[i].ValidatorAddress < av[j].ValidatorAddress
}

func (av AllowListedVals) Swap(i, j int) {
	av[i], av[j] = av[j], av[i]
}

// NewHostAccountDelegation returns new HostAccountDelegation
func NewHostAccountDelegation(validatorAddress string, amount sdk.Coin) HostAccountDelegation {
	return HostAccountDelegation{
		ValidatorAddress: validatorAddress,
		Amount:           amount,
	}
}

func (h HostAccountDelegations) Len() int {
	return len(h)
}

func (h HostAccountDelegations) Less(i, j int) bool {
	return h[i].ValidatorAddress < h[j].ValidatorAddress
}

func (h HostAccountDelegations) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// NewHostChainRewardAddress returns new HostChainRewardAddress
func NewHostChainRewardAddress(address string) HostChainRewardAddress {
	return HostChainRewardAddress{
		Address: address,
	}
}

// ValAddressFromBech32 creates a ValAddress from a Bech32 string.
func ValAddressFromBech32(address string, bech32Prefix string) (addr sdk.ValAddress, err error) {
	if len(strings.TrimSpace(address)) == 0 {
		return sdk.ValAddress{}, errors.New("empty address string is not allowed")
	}

	bz, err := sdk.GetFromBech32(address, bech32Prefix)
	if err != nil {
		return nil, err
	}

	err = sdk.VerifyAddressFormat(bz)
	if err != nil {
		return nil, err
	}

	return bz, nil
}

func Bech32FromValAddress(valAddr sdk.ValAddress, bech32Prefix string) (string, error) {
	if valAddr.Empty() {
		return "", sdkerrors.ErrInvalidAddress
	}
	bech32Addr, err := bech32.ConvertAndEncode(bech32Prefix, valAddr)
	if err != nil {
		return "", err
	}
	return bech32Addr, nil
}

// TotalDelegations gives the amount of total delegations on Host Chain.
func (ds DelegationState) TotalDelegations(denom string) sdk.Coin {
	total := sdk.NewCoin(denom, sdk.ZeroInt())

	for _, val := range ds.HostAccountDelegations {
		if val.Amount.Denom == denom {
			total.Amount = total.Amount.Add(val.Amount.Amount)
		}
	}
	return total
}

// NewDelegatorUnbondingEpochEntry returns new DelegatorUnbondingEpochEntry
func NewDelegatorUnbondingEpochEntry(delegatorAddress string, epochNumber int64, amount sdk.Coin) DelegatorUnbondingEpochEntry {
	return DelegatorUnbondingEpochEntry{
		DelegatorAddress: delegatorAddress,
		EpochNumber:      epochNumber,
		Amount:           amount,
	}
}

// GetUnbondingEpochCValue returns the calculated c value from the UnbondingEpochCValue struct entries.
func (uec *UnbondingEpochCValue) GetUnbondingEpochCValue() sdk.Dec {
	return sdk.NewDecFromInt(uec.STKBurn.Amount).Quo(sdk.NewDecFromInt(uec.AmountUnbonded.Amount))
}

// CurrentUnbondingEpoch computes and returns current unbonding epoch to the next nearest multiple
// 4 UndelegationEpochNumberFactor
func CurrentUnbondingEpoch(epochNumber int64) int64 {
	if epochNumber%UndelegationEpochNumberFactor == 0 {
		return epochNumber
	}
	return epochNumber + UndelegationEpochNumberFactor - epochNumber%UndelegationEpochNumberFactor
}

// PreviousUnbondingEpoch computes and returns previous unbonding epoch to the previous nearest
// multiple of UndelegationEpochNumberFactor
func PreviousUnbondingEpoch(epochNumber int64) int64 {
	if epochNumber%UndelegationEpochNumberFactor == 0 {
		return epochNumber - UndelegationEpochNumberFactor
	}
	return epochNumber - epochNumber%UndelegationEpochNumberFactor
}

// DelegatorAccountPortID returns  delegator account port ID
func (hostAccounts *HostAccounts) DelegatorAccountPortID() string {
	delegatorAccountPortID, err := icatypes.NewControllerPortID(hostAccounts.DelegatorAccountOwnerID)
	if err != nil {
		panic(err)
	}
	return delegatorAccountPortID
}

// RewardsAccountPortID returns rewards account port ID
func (hostAccounts *HostAccounts) RewardsAccountPortID() string {
	rewardsAccountPortID, err := icatypes.NewControllerPortID(hostAccounts.RewardsAccountOwnerID)
	if err != nil {
		panic(err)
	}
	return rewardsAccountPortID
}

// Validate returns error if contents do not pass the expected checks
func (hostAccounts *HostAccounts) Validate() error {
	if hostAccounts.RewardsAccountOwnerID == "" || hostAccounts.DelegatorAccountOwnerID == "" {
		return ErrInvalidHostAccountOwnerIDs
	}
	return nil
}

func (gstakeParams *PstakeParams) Validate() error {
	_, err := sdk.AccAddressFromBech32(gstakeParams.PstakeFeeAddress)
	if err != nil {
		return err
	}

	if gstakeParams.PstakeDepositFee.IsNegative() || gstakeParams.PstakeDepositFee.GTE(MaxPstakeDepositFee) {
		return errorsmod.Wrapf(ErrInvalidFee, "gstake deposit fee must be between %s and %s", sdk.ZeroDec(), MaxPstakeDepositFee)
	}

	if gstakeParams.PstakeRestakeFee.IsNegative() || gstakeParams.PstakeRestakeFee.GTE(MaxPstakeRestakeFee) {
		return errorsmod.Wrapf(ErrInvalidFee, "gstake restake fee must be between %s and %s", sdk.ZeroDec(), MaxPstakeRestakeFee)
	}

	if gstakeParams.PstakeUnstakeFee.IsNegative() || gstakeParams.PstakeUnstakeFee.GTE(MaxPstakeUnstakeFee) {
		return errorsmod.Wrapf(ErrInvalidFee, "gstake unstake fee must be between %s and %s", sdk.ZeroDec(), MaxPstakeUnstakeFee)
	}

	if gstakeParams.PstakeRedemptionFee.IsNegative() || gstakeParams.PstakeRedemptionFee.GTE(MaxPstakeRedemptionFee) {
		return errorsmod.Wrapf(ErrInvalidFee, "gstake redemption fee must be between %s and %s", sdk.ZeroDec(), MaxPstakeRedemptionFee)
	}
	return nil
}

func ConvertMintDenomToBaseDenom(mintDenom string) (string, error) {
	denomSplit := strings.Split(mintDenom, "/")

	if denomSplit[0] != LiquidStakedDenomPrefix || len(denomSplit) != 2 {
		return "", ErrInvalidMintDenom
	}
	return denomSplit[1], nil

}

func ConvertBaseDenomToMintDenom(baseDenom string) string {
	return fmt.Sprintf("%s/%s", LiquidStakedDenomPrefix, baseDenom)
}
