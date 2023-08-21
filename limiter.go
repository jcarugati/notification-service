package main

import (
	"github.com/patrickmn/go-cache"
)

// LocalRateLimiter is a local rate limiter that uses a cache to limit requests.
type LocalRateLimiter struct {
	c cache.Cache
}

// Allow checks if a request is allowed based on a key and a given ttl.
// It returns false if the key is found in the cache (request not allowed)
// and true otherwise (request allowed).
func (lrl *LocalRateLimiter) Allow(key string, rule *Rule) bool {
	value, found := lrl.c.Get(key)
	if !found {
		value = 0
	}

	before, ok := value.(int)
	if !ok {
		return false
	}

	after := before + 1

	if after > rule.MaxAttempts {
		return false
	}

	lrl.c.Set(key, after, rule.TTL)

	return true
}

// NewLocalRateLimiter initializes a new LocalRateLimiter with default cache settings.
func NewLocalRateLimiter() *LocalRateLimiter {
	return &LocalRateLimiter{c: *cache.New(cache.NoExpiration, cache.NoExpiration)}
}
