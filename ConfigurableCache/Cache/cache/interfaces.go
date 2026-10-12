// Package cache is a cache with pluggable storage and a pluggable eviction
// policy.
//
// The design has three parts:
//
//	Storage         WHERE items live        (interface; in-memory implemented)
//	EvictionPolicy  WHO gets removed first  (interface; LRU implemented)
//	Cache           coordinates the two, and owns the rules that are the same
//	                for every combination: max capacity, expiry, locking
//
// To add LFU or a file-system store later, you write a new type that
// satisfies the interface. The Cache code does not change at all.
package cache

import "time"

// ---------------------------------------------------------------------------
// THE TWO INTERFACES (the "plug sockets")
// ---------------------------------------------------------------------------
//
// In Go a type satisfies an interface automatically, just by having the right
// methods. There is no "implements" keyword.

// Item is what the cache hands to storage: the value plus its expiry time.
// Expiry is stored WITH the item so that a persistent store (a file, say)
// would remember it across restarts.
type Item[V any] struct {
	Value     V
	ExpiresAt time.Time
}

func (i Item[V]) expired(now time.Time) bool {
	return !i.ExpiresAt.IsZero() && now.After(i.ExpiresAt)
}

// Storage only knows how to keep and fetch items. It knows nothing about
// eviction or expiry
//
// Methods return an error because some store can fail (disk full, file
// missing) The in-memory one never fails, but the interface has to allow
// for the ones that do
type Storage[K comparable, V any] interface {
	Get(key K) (item Item[V], found bool, err error)
	Set(key K, item Item[V]) error
	Delete(key K) error
	Len() int
}

// EvictionPolicy only tracks and decides which one should go next.
// It never sees the values
//
// The cache tells the policy what happened (insert/access/remove), and
// asks it for a victim when cache is full.
type EvictionPolicy[K comparable] interface {
	OnInsert(key K) // a new key was added
	OnAccess(key K) // an existing key was read or updated
	OnRemove(key K) // a key was deleted or expired
	// Evict picks the key to remove and forgets it
	// It returns false if the policy is tracking no keys
	Evict() (key K, ok bool)
}
