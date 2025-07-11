package v11

import (
	store "github.com/cosmos/cosmos-sdk/store/types"
	liquidstaketypes "github.com/persistenceOne/pstake-native/v3/x/liquidstake/types"

	"github.com/persistenceOne/persistenceCore/v12/app/upgrades"
)

const (
	// UpgradeName defines the on-chain upgrade name.
	UpgradeName = "v11"
)

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{
			liquidstaketypes.StoreKey,
		},
	},
}
