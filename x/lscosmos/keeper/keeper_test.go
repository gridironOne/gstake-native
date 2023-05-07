package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank/testutil"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	ibctransfertypes "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"
	ibctesting "github.com/cosmos/ibc-go/v6/testing"
	"github.com/stretchr/testify/suite"

	"github.com/gridironOne/gstake-native/v2/app"
	"github.com/gridironOne/gstake-native/v2/app/helpers"
	"github.com/gridironOne/gstake-native/v2/x/lscosmos/types"
)

var (
	allowListedValidators = types.AllowListedValidators{
		AllowListedValidators: []types.AllowListedValidator{
			{
				ValidatorAddress: "cosmosvaloper1hcqg5wj9t42zawqkqucs7la85ffyv08le09ljt",
				TargetWeight:     sdk.NewDecWithPrec(33, 2),
			},
			{
				ValidatorAddress: "cosmosvaloper1lcck2cxh7dzgkrfk53kysg9ktdrsjj6jfwlnm2",
				TargetWeight:     sdk.NewDecWithPrec(33, 2),
			},
			{
				ValidatorAddress: "cosmosvaloper10khgeppewe4rgfrcy809r9h00aquwxxxgwgwa5",
				TargetWeight:     sdk.NewDecWithPrec(34, 2),
			},
		},
	}
	ChainID          = "cosmoshub-4"
	ConnectionID     = "connection-0"
	TransferChannel  = "channel-0"
	TransferPort     = "transfer"
	BaseDenom        = "uatom"
	MintDenom        = "stk/uatom"
	MinDeposit       = sdk.NewInt(5)
	GstakeFeeAddress = "gridiron1pss7nxeh3f9md2vuxku8q99femnwdjtcpe9ky9"
)

func init() {
	ibctesting.DefaultTestingAppInit = helpers.SetupTestingApp
}

type IntegrationTestSuite struct {
	suite.Suite

	app        *app.GstakeApp
	ctx        sdk.Context
	govHandler govtypes.Handler

	coordinator *ibctesting.Coordinator
	chainA      *ibctesting.TestChain
	chainB      *ibctesting.TestChain
	path        *ibctesting.Path
}

func newGstakeAppPath(chainA, chainB *ibctesting.TestChain) *ibctesting.Path {
	path := ibctesting.NewPath(chainA, chainB)
	path.EndpointA.ChannelConfig.PortID = ibctesting.TransferPort
	path.EndpointB.ChannelConfig.PortID = ibctesting.TransferPort

	return path
}

func GetGstakeApp(chain *ibctesting.TestChain) *app.GstakeApp {
	app1, ok := chain.App.(*app.GstakeApp)
	if !ok {
		panic("not gstake app")
	}

	return app1
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (suite *IntegrationTestSuite) SetupTest() {
	_, gstakeApp, ctx := helpers.CreateTestApp(suite.T())

	keeper := gstakeApp.LSCosmosKeeper

	params := types.DefaultParams()
	keeper.SetParams(ctx, params)

	suite.app = &gstakeApp
	suite.ctx = ctx

	suite.coordinator = ibctesting.NewCoordinator(suite.T(), 2)
	suite.chainA = suite.coordinator.GetChain(ibctesting.GetChainID(1))
	suite.chainB = suite.coordinator.GetChain(ibctesting.GetChainID(2))

	suite.path = newGstakeAppPath(suite.chainA, suite.chainB)
	suite.coordinator.SetupConnections(suite.path)

	// set host chain params
	depositFee, err := sdk.NewDecFromStr("0.01")
	suite.NoError(err)

	restakeFee, err := sdk.NewDecFromStr("0.02")
	suite.NoError(err)

	unstakeFee, err := sdk.NewDecFromStr("0.03")
	suite.NoError(err)

	redemptionFee, err := sdk.NewDecFromStr("0.03")
	suite.NoError(err)

	hostChainParams := types.NewHostChainParams(
		ChainID,
		ConnectionID,
		TransferChannel,
		TransferPort,
		BaseDenom,
		MintDenom,
		GstakeFeeAddress,
		MinDeposit,
		depositFee,
		restakeFee,
		unstakeFee,
		redemptionFee,
	)
	suite.app.LSCosmosKeeper.SetHostChainParams(suite.ctx, hostChainParams)
	suite.app.LSCosmosKeeper.SetHostAccounts(suite.ctx, types.HostAccounts{
		DelegatorAccountOwnerID: "Del_acc",
		RewardsAccountOwnerID:   "Rew_acc",
	})
	suite.app.LSCosmosKeeper.SetAllowListedValidators(ctx, allowListedValidators)
}

func (suite *IntegrationTestSuite) TestMintToken() {
	gstakeApp, ctx := suite.app, suite.ctx
	testParams := gstakeApp.LSCosmosKeeper.GetHostChainParams(ctx)

	ibcDenom := ibctransfertypes.GetPrefixedDenom(testParams.TransferPort, testParams.TransferChannel, testParams.BaseDenom)
	balanceOfIbcToken := sdk.NewInt64Coin(ibcDenom, 100)
	mintAmountDec := sdk.NewDecFromInt(balanceOfIbcToken.Amount).Mul(gstakeApp.LSCosmosKeeper.GetCValue(ctx))
	toBeMintedTokens, _ := sdk.NewDecCoinFromDec(testParams.MintDenom, mintAmountDec).TruncateDecimal()

	addr := sdk.AccAddress("addr________________")
	acc := gstakeApp.AccountKeeper.NewAccountWithAddress(ctx, addr)
	gstakeApp.AccountKeeper.SetAccount(ctx, acc)
	suite.Require().NoError(testutil.FundAccount(gstakeApp.BankKeeper, ctx, addr, sdk.NewCoins(balanceOfIbcToken)))

	suite.Require().NoError(gstakeApp.LSCosmosKeeper.MintTokens(ctx, toBeMintedTokens, addr))

	currBalance := gstakeApp.BankKeeper.GetBalance(ctx, addr, testParams.MintDenom)

	suite.Require().Equal(toBeMintedTokens, currBalance)
}
