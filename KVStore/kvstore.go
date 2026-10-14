package kvstore

import (
	"errors"
	"sync"
)


var (
	ErrNoTransaction = errors.New("no active transaction")
)
// entry represents a value in a transaction
// we use a deleted flag (tombstone) to track keys deleted during a transaction.
type entry struct {
	value string
	deleted bool
}

// transaction represents a single change set in the stack
type transaction struct {
	changes map[string]entry
}

// DB represents our key-value store with transaction support.
type DB struct {
	mu 		sync.RWMutex
	store 	map[string]string
	txs 	[]*transaction	// Stack of nested transactions
}

// NewDB initializes and returns a new DB instance.
func NewDB() *DB {
	return &DB{
		store: make(map[string]string),
		txs: make([]*transaction, 0),
	}
}

func (db *DB) Set(key, value string) {
	db.mu.Lock()
	defer db.mu.Unlock()

	if len(db.txs) > 0 {
		// Apply to the current (topmost) transaction
		topTx := db.txs[len(db.txs)-1]
		topTx.changes[key] = entry{value: value, deleted: false}
		return
	}

	// Apply directly to the base store
	db.store[key] = value
}

func (db *DB) Get(key string) (string, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	// 1. Search top-down through the transaction stack
	for i := len(db.txs)-1; i >= 0; i-- {
		tx := db.txs[i]
		if e, exists := tx.changes[key]; exists {
			if e.deleted {
				return "", false
			}
			return e.value, true
		}
	}

	// 2. Search base store if not found in any transaction
	val, exists := db.store[key]
	return val, exists
}

func (db *DB) Delete(key string) {
	db.mu.Lock()
	defer db.mu.Unlock()

	if len(db.txs) > 0 {
		// Mark as deleted in current transaction (Tombstone)
		topTx := db.txs[len(db.txs)-1]
		topTx.changes[key] = entry{deleted: true}
		return
	}

	delete(db.store, key)
}

func (db *DB) Begin() {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.txs = append(db.txs, &transaction{
		changes: make(map[string]entry),
	})
}

// Commit merges the topmost transaction into its parent (or base store)
func (db *DB) Commit() error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if len(db.txs) == 0 {
		return ErrNoTransaction
	}

	//Pop the topmost transaction
	topIdx := len(db.txs) - 1
	topTx := db.txs[topIdx]
	db.txs = db.txs[:topIdx]
	
	// Merge logic
	if len(db.txs) > 0 {
		// Merge into parent transaction
		parentTx := db.txs[len(db.txs)-1]
		for k,e := range topTx.changes {
			parentTx.changes[k] = e
		}
	} else {
		for k, e := range topTx.changes {
			if e.deleted {
				delete(db.store, k)
			} else {
				db.store[k] = e.value
			}
		}
	}
	return nil
}

func (db *DB) Rollback() error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if len(db.txs) == 0 {
		return ErrNoTransaction
	}

	// Pop and discard the topmost transaction
	db.txs = db.txs[:len(db.txs)-1]
	return nil
}