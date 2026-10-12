package cache

import (
	"errors"
	"sync"
	"time"
)

// ErrNotFound is a "sentinel error": a named value callers can compare
// against with errors.Is(err, ErrNotFound).
var ErrNotFound = errors.New("cache: key not found")

// Config is how the client configures the cache.
//
// (The first version used functional options. With generic plug-ins a plain
// struct reads better, because Go can work out K and V from the struct and
// you don't have to repeat them on every option.)
//
// Every field is optional. Leave one out and you get the default.
type Config[K comparable, V any] struct {
	Maxsize    int
	DefaultTTL time.Duration
	Storage    Storage[K, V]
	Policy     EvictionPolicy[K]
}

type Cache[K comparable, V any] struct {
	mu         sync.Mutex
	maxSize    int
	defaultTTL time.Duration
	storage    Storage[K, V]
	policy     EvictionPolicy[K]
}

func New[K comparable, V any](cfg Config[K, V]) *Cache[K, V] {
	if cfg.Storage == nil {
		cfg.Storage = NewMemoryStorage[K, V]()
	}
	if cfg.Policy == nil {
		cfg.Policy = NewLRU[K]()
	}
	return &Cache[K, V]{
		maxSize:    cfg.Maxsize,
		defaultTTL: cfg.DefaultTTL,
		storage:    cfg.Storage,
		policy:     cfg.Policy,
	}
}

// Set stores a value using the cache's default TTL.
func (c *Cache[K, V]) Set(key K, value V) error {
	return c.SetWithTTL(key, value, c.defaultTTL)
}

func (c *Cache[K, V]) SetWithTTL(key K, value V, ttl time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	item := Item[V]{Value: value}
	if ttl > 0 {
		item.ExpiresAt = time.Now().Add(ttl)
	}

	_, exists, err := c.storage.Get(key)

	if err != nil {
		return err
	}

	if !exists && c.maxSize > 0 && c.storage.Len() >= c.maxSize {
		if victim, ok := c.policy.Evict(); ok {
			if err := c.storage.Delete(victim); err != nil {
				return err
			}
		}
	}

	if err := c.storage.Set(key, item); err != nil {
		return err
	}

	if exists {
		c.policy.OnAccess(key)
	} else {
		c.policy.OnInsert(key)
	}
	return nil
}

func (c *Cache[K, V]) Get(key K) (V, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var zero V

	item, found, err := c.storage.Get(key)
	if err != nil {
		return zero, err
	}

	if !found {
		return zero, ErrNotFound
	}

	if item.expired(time.Now()) {
		if err := c.remove(key); err != nil {
			return zero, err
		}
		return zero, ErrNotFound
	}

	c.policy.OnAccess(key)
	return item.Value, nil
}

func (c *Cache[K, V]) Delete(key K) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if err := c.remove(key); err != nil {
		return err
	}

	return nil
}

func (c *Cache[K, V]) Len() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.storage.Len()
}

func (c *Cache[K, V]) remove(key K) error {
	if err := c.storage.Delete(key); err != nil {
		return err
	}
	c.policy.OnRemove(key)
	return nil
}
