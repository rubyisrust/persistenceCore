package v11_1_0

import (
	store "github.com/cosmos/cosmos-sdk/store/types"
	ratesynctypes "github.com/persistenceOne/pstake-native/v3/x/ratesync/types"

	"github.com/persistenceOne/persistenceCore/v12/app/upgrades"
)

const (
	// UpgradeName defines the on-chain upgrade name.
	UpgradeName = "v11.1.0"
)

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{ratesynctypes.StoreKey},
	},
}
