package v004

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	tokenfactorykeeper "github.com/cosmostation/mintstation/x/tokenfactory/keeper"
	tokenfactorytypes "github.com/cosmostation/mintstation/x/tokenfactory/types"
)

var (
	UpgradeName = "v0.0.4"
)

// CreateUpgradeHandler creates an SDK upgrade handler for v003
func CreateUpgradeHandler(
	mm *module.Manager,
	tokenFactoryKeeper tokenfactorykeeper.Keeper,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		vm, err := mm.RunMigrations(ctx, configurator, vm)

		if err := tokenFactoryKeeper.SetParams(ctx, tokenfactorytypes.Params{
			DenomCreationFee:        nil,
			DenomCreationGasConsume: 50_000,
		}); err != nil {
			return nil, err
		}

		return vm, err
	}
}
