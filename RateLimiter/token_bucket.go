package ratelimiter

import (
	"sync"
	"time"
)

type Clock func() time.Time

type bucket struct {
	tokens float64
	last time.Time
}

type TokenBucket struct {
	mu sync.Mutex
	buckets map[string]*bucket
	capacity float64
	rate float64
	now Clock
}

func NewTokenBucket(capacity float64, rate float64, now Clock) *TokenBucket {
	if now == nil {
		now = time.Now
	}

	return &TokenBucket{
		buckets: 	make(map[string]*bucket),
		capacity: 	capacity,
		rate: 		rate,
		now: 		now,
	}
}

var _ Limiter = (*TokenBucket)(nil)

func (tb *TokenBucket) Allow(userId string) bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	now := tb.now()
	b, ok := tb.buckets[userId]

	if !ok {
		b = &bucket{tokens: tb.capacity, last: now}
		tb.buckets[userId] = b
	}

	// Refill based on time passed since last call
	elapsed := now.Sub(b.last).Seconds()
	b.tokens = min(tb.capacity, b.tokens + (elapsed*tb.rate))
	b.last = now

	// Spend token if there is one
	if b.tokens >= 1 {
		b.tokens--
		return true
	}
	return false
}