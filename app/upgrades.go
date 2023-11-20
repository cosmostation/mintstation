package app

import (
	"fmt"

	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	v003 "github.com/cosmostation/mintstation/app/upgrades/v003"
	mnsmoduletypes "github.com/cosmostation/mintstation/x/mns/types"
)

func (app *App) setupUpgradeHandlers(appOpts servertypes.AppOptions) {
	// v003 upgrade handler
	app.UpgradeKeeper.SetUpgradeHandler(
		v003.UpgradeName,
		v003.CreateUpgradeHandler(
			app.mm,
			app.configurator,
		),
	)

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Errorf("Failed to read upgrade info from disk: %w", err))
	}

	if app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		return
	}

	var storeUpgrades *storetypes.StoreUpgrades

	switch upgradeInfo.Name {
	case "v0.0.3":
		storeUpgrades = &storetypes.StoreUpgrades{
			Added: []string{mnsmoduletypes.StoreKey},
		}
	}

	if storeUpgrades != nil {
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, storeUpgrades))
	}
}
