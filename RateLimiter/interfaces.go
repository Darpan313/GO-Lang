package ratelimiter

type Limiter interface {
	Allow(userId string) bool
}
