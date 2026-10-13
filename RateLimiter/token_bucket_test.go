package ratelimiter

import (
	"testing"
	"time"
)

type fakeClock struct {
	t time.Time
}

func (f *fakeClock) Now() time.Time {
	return f.t
}
func (f *fakeClock) Advance(d time.Duration) {
	f.t = f.t.Add(d)
}

func newTestLimiter(capacity, rate float64) (*TokenBucket, *fakeClock) {
	fc := &fakeClock{t: time.Unix(0,0)}
	return NewTokenBucket(capacity, rate, fc.Now), fc
}

func TestBurst(t *testing.T) {
	const capacity = 3
	lim, _ := newTestLimiter(capacity, 1)
	for i := range capacity {
		if !lim.Allow("alice") {
			t.Fatalf("request %d: got denied, want allowed", i)
		}
	}
	if lim.Allow("alice") {
		t.Errorf("request %d: got allowed, want denied", capacity+1)
	}
}

func TestRefillAfterAdvanceWithRate1(t *testing.T) {
	const capacity = 3
	lim, clock := newTestLimiter(capacity, 1)
	for i := range capacity {
		if !lim.Allow("alice") {
			t.Fatalf("request %d: got denied, want allowed", i)
		}
	}
	
	clock.Advance(time.Second)

	if !lim.Allow("alice") {
		t.Fatalf("request %d: got denied, want allowed", capacity+1)
	}

	if lim.Allow("alice") {
		t.Errorf("request %d: got allowed, want denied", capacity+2)
	}
}

func TestWaitingALongTimeNotOverfill(t *testing.T) {
	
}

func TestUserDoNotAffectEachOther(t *testing.T) {
	
}

func TestConcurrentRequestsOnlyCapacityAllowed(t *testing.T) {
	
}