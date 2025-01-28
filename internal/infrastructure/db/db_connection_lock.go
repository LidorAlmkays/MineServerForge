package db

import (
	"sync"

	"github.com/LidorAlmkays/MineServerForge/pkg/utils/enums"
)

// LockManager dynamically manages mutexes based on keys.
type dbConnectionLockManager struct {
	mutexes map[enums.DbType]*sync.Mutex
	mu      sync.Mutex // protects access to the mutexes map
}

func newDbConnectionLockManager() *dbConnectionLockManager {
	return &dbConnectionLockManager{mutexes: map[enums.DbType]*sync.Mutex{}}
}

// Lock locks the mutex for the given key.
func (lm *dbConnectionLockManager) Lock(key enums.DbType) {
	lm.mu.Lock()
	if _, exists := lm.mutexes[key]; !exists {
		lm.mutexes[key] = &sync.Mutex{}
	}
	mutex := lm.mutexes[key]
	lm.mu.Unlock()

	mutex.Lock()
}

// Unlock unlocks the mutex for the given key.
func (lm *dbConnectionLockManager) Unlock(key enums.DbType) {
	lm.mu.Lock()
	mutex, exists := lm.mutexes[key]
	lm.mu.Unlock()

	if exists {
		mutex.Unlock()
	}
}
