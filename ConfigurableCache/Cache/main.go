package main

// Demo of the cache package.
//
// Run it with:   go run .

import (
	"errors"
	"fmt"
	"time"

	"configurablecache/cache"
)

func main() {
	c := cache.New(cache.Config[string, int]{
		Maxsize:    2,
		DefaultTTL: 100 * time.Millisecond,
		Storage:    cache.NewMemoryStorage[string, int](),
		Policy:     cache.NewLRU[string](),
	})

	c.Set("a", 1)
	c.Set("b", 2)

	c.Get("a")
	c.Set("c", 3)

	_, err := c.Get("b")
	fmt.Println("b evicted?", errors.Is(err, cache.ErrNotFound)) // true

	v, err := c.Get("a")
	fmt.Println("a =", v, "err =", err) // 1 <nil>

	// A per-key TTL overrides the default.
	c.SetWithTTL("long", 99, time.Hour) // evicts "c" (least recently used)

	time.Sleep(150 * time.Millisecond) // wait past the default TTL

	_, err = c.Get("a")
	fmt.Println("a expired?", errors.Is(err, cache.ErrNotFound)) // true

	v, err = c.Get("long")
	fmt.Println("long =", v, "err =", err) // 99 <nil>

	// Leave everything out and you get unlimited, in-memory, LRU, no expiry.
	names := cache.New(cache.Config[int, string]{})
	names.Set(1, "Ada")
	name, _ := names.Get(1)
	fmt.Println("name:", name)
}
