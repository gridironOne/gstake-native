package lscosmos_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/stretchr/testify/suite"

	"github.com/gridironOne/gstake-native/app"
	"github.com/gridironOne/gstake-native/app/helpers"
	"github.com/gridironOne/gstake-native/x/lscosmos"
)

type HandlerTestSuite struct {
	suite.Suite

	app        *app.GstakeApp
	ctx        sdk.Context
	govHandler govtypes.Handler
}

func (suite *HandlerTestSuite) SetupTest() {
	_, gstakeApp, ctx := helpers.CreateTestApp(suite.T())
	suite.app = &gstakeApp
	suite.ctx = ctx
	suite.govHandler = lscosmos.NewLSCosmosProposalHandler(suite.app.LSCosmosKeeper)
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}
