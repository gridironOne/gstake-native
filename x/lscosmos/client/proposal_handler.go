package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/gridironOne/gstake-native/v2/x/lscosmos/client/cli"
)

var (
	MinDepositAndFeeChangeProposalHandler      = govclient.NewProposalHandler(cli.NewMinDepositAndFeeChangeCmd)
	GstakeFeeAddressChangeProposalHandler      = govclient.NewProposalHandler(cli.NewGstakeFeeAddressChangeCmd)
	AllowListValidatorSetChangeProposalHandler = govclient.NewProposalHandler(cli.NewAllowListedValidatorSetChangeProposalCmd)
)
