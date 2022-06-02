package v2_test

import (
	"fmt"
	"testing"

	"github.com/Karan-3108/ethermint/app"
	"github.com/Karan-3108/ethermint/encoding"
	v2 "github.com/Karan-3108/ethermint/x/evm/migrations/v2"
	"github.com/Karan-3108/ethermint/x/evm/types"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"
)

func TestAddMinGasMultiplierParam(t *testing.T) {
	encCfg := encoding.MakeConfig(app.ModuleBasics)
	erc20Key := sdk.NewKVStoreKey(types.StoreKey)
	tErc20Key := sdk.NewTransientStoreKey(fmt.Sprintf("%s_test", types.StoreKey))
	ctx := testutil.DefaultContext(erc20Key, tErc20Key)
	paramStore := paramtypes.NewSubspace(
		encCfg.Marshaler, encCfg.Amino, erc20Key, tErc20Key, "erc20",
	)
	paramStore = paramStore.WithKeyTable(types.ParamKeyTable())
	require.True(t, paramStore.HasKeyTable())

	// check no param
	require.False(t, paramStore.Has(ctx, types.ParamStoreKeyMinGasMultiplier))

	// Run migrations
	err := v2.AddMinGasMultiplierParam(ctx, &paramStore)
	require.NoError(t, err)

	// Make sure the params are set
	require.True(t, paramStore.Has(ctx, types.ParamStoreKeyMinGasMultiplier))

	var minGasMultiplier sdk.Dec

	// Make sure the new params are set
	require.NotPanics(t, func() {
		paramStore.Get(ctx, types.ParamStoreKeyMinGasMultiplier, &minGasMultiplier)
	})

	// check the params is up
	require.Equal(t, minGasMultiplier, types.DefaultMinGasMultiplier)
}
