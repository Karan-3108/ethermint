package keeper_test

import "github.com/Karan-3108/ambinet/v4/x/erc20/types"

func (suite *KeeperTestSuite) TestParams() {
	params := suite.app.Erc20Keeper.GetParams(suite.ctx)
	suite.Require().Equal(types.DefaultParams(), params)
	params.EnableErc20 = false
	suite.app.Erc20Keeper.SetParams(suite.ctx, params)
	newParams := suite.app.Erc20Keeper.GetParams(suite.ctx)
	suite.Require().Equal(newParams, params)
}
