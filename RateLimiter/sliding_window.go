package ratelimiter

import (
	"sync"
	"time"
)

type window struct {
	start time.Time
	currCount int
	prevCount int
}

type SlidingWindow struct {
	mu sync.Mutex
	windows map[string]*window
	limit int
	size time.Duration
	now Clock
}

func NewSlidingWindow(limit int, size time.Duration, now Clock) *SlidingWindow {
	if now == nil {
		now = time.Now
	}

	return &SlidingWindow{
		windows: make(map[string]*window)
		limit: limit,
		size: size,
		now: now,
	}
}

var _ Limiter = (*SlidingWindow)(nil)

func(sw *SlidingWindow) Allow(userId string) bool {
	sw.mu.Lock()
	defer sw.mu.Unlock()

	now := sw.now()
	start := now.Truncate(sw.size)

	w, ok := sw.windows[userId]
	if !ok {
		w = &window{start: start}
		sw.windows[userId] = w
	}

	if start != w.start {
		if start.Sub(w.start) == sw.size {
			w.prevCount = w.currCount	// adjacent window: carry it over
		} else {
			w.prevCount = 0 // skipped one or more windows: nothing overlaps
		}
		w.currCount = 0
		w.start = start
	}

	elapsed := now.Sub(w.start).Seconds()/sw.size.Seconds()
	estimate := float64(w.prevCount)*(1-elapsed) + float64(w.currCount)

	if estimate < float64(sw.limit) {
		w.currCount++
		return true
	}
	return false
}