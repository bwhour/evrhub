package store

import (
	dbm "github.com/tendermint/tm-db"

	"github.com/Evrynetlabs/evrsdk/store/cache"
	"github.com/Evrynetlabs/evrsdk/store/rootmulti"
	"github.com/Evrynetlabs/evrsdk/store/types"
)

func NewCommitMultiStore(db dbm.DB) types.CommitMultiStore {
	return rootmulti.NewStore(db)
}

func NewCommitKVStoreCacheManager() types.MultiStorePersistentCache {
	return cache.NewCommitKVStoreCacheManager(cache.DefaultCommitKVStoreCacheSize)
}
