package v2

import (
	"github.com/Karan-3108/ethermint/x/evm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// AddMinGasMultiplierParam updates the module parameter MinGasMultiplier to 0.5
func AddMinGasMultiplierParam(ctx sdk.Context, paramStore *paramtypes.Subspace) error {
	if !paramStore.HasKeyTable() {
		ps := paramStore.WithKeyTable(types.ParamKeyTable())
		paramStore = &ps
	}

	paramStore.Set(ctx, types.ParamStoreKeyMinGasMultiplier, types.DefaultMinGasMultiplier)
	return nil
}
