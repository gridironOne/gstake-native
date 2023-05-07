<!--
order: 2
-->

# Change Min Deposit and Fee Proposal

This proposal takes the following parameters into account :

- `MinDeposit` : of type sdk.Int
- `GstakeDepositFee` : of type sdk.Dec
- `GstakeRestakeFee` : of type sdk.Dec
- `GstakeUnstakeFee` : of type sdk.Dec
- `GstakeRedemptionFee` : of type sdk.Dec

Example proposal :

```json
{
  "title": "min-deposit and fee change proposal",
  "description": "this proposal changes min-deposit and protocol fee on chain",
  "min_deposit": "5",
  "gstake_deposit_fee": "0.1",
  "gstake_restake_fee": "0.1",
  "gstake_unstake_fee": "0.1",
  "gstake_redemption_fee": "0.1",
  "deposit": "1000000stake"
}
```

Sample command to submit proposal :

```
$ $BIN_NAME tx gov submit-proposal gstake-lscosmos-min-deposit-and-fee-change  <path/to/proposal.json> --from <key_or_address> --fees <1000stake> --gas <200000>
```

It used when there is a need to change all the different types of fees and minimum deposit needed by user to liquid
stake.

# Change Gstake Fee Address Proposal

This proposal take the following prameters into account :

- `GstakeFeeAddress` : It is the address in which gstake protocol fee is sent.

Example proposal :

```json
{
  "title": "change gstake fee address",
  "description": "this proposal changes gstake fee address in the chain",
  "gstake_fee_address": "gridiron1pss7nxeh3f9md2vuxku8q99femnwdjtcpe9ky9",
  "deposit": "10000000stake"
}
```

Sample command to submit proposal :

```
$ $BIN_NAME tx gov submit-proposal gstake-lscosmos-change-gstake-fee-address <path/to/proposal.json> --from <key_or_address> --fees <1000stake> --gas <200000>
```

It is used to change the gstake fee address if old one is not needed anymore.

# Change Allow Listed Validators Proposal

This proposal take the following prameters into account :

- `AllowListedValidators` : List of validators and their corresponding weights.

Example proposal :

```json
{
  "title": "change gstake fee address",
  "description": "this proposal changes gstake fee address in the chain",
  "allow_listed_validators": {
    "allow_listed_validators": [
      {
        "validator_address": "cosmosvaloper1hcqg5wj9t42zawqkqucs7la85ffyv08le09ljt",
        "target_weight": "1"
      }
    ]
  },
  "deposit": "100000stake"
}
```

Sample command to submit proposal :

```
$ $BIN_NAME tx gov submit-proposal gstake-lscosmos-change-allow-listed-validator-set <path/to/proposal.json> --from <key_or_address> --fees <1000stake> --gas <200000>
```

It is to change the validator set of lscsomos module if in case the old validator set becomes stale.


