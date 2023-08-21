package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewLocalRateLimiter(t *testing.T) {
	limiter := NewLocalRateLimiter()
	assert.NotNil(t, limiter)
}

func TestAllowFirstTime(t *testing.T) {
	limiter := NewLocalRateLimiter()

	// First call to Allow for a given key should return true.
	assert.True(t, limiter.Allow("key1", time.Second))
}

func TestDisallowWithinTTL(t *testing.T) {
	limiter := NewLocalRateLimiter()
	ttl := time.Second

	// Set an item with TTL.
	limiter.Allow("key1", ttl)

	// Within TTL, Allow should return false.
	assert.False(t, limiter.Allow("key1", ttl))
}

func TestAllowAfterTTLExpires(t *testing.T) {
	limiter := NewLocalRateLimiter()
	ttl := 100 * time.Millisecond

	// Set an item with TTL.
	limiter.Allow("key1", ttl)

	// Wait for the TTL to expire.
	time.Sleep(ttl + 50*time.Millisecond)

	// After TTL expires, Allow should return true.
	assert.True(t, limiter.Allow("key1", ttl))
}

func TestDifferentKeys(t *testing.T) {
	limiter := NewLocalRateLimiter()
	ttl := time.Second

	// Set an item with TTL for key1.
	limiter.Allow("key1", ttl)

	// Different key should not be affected by the TTL of key1.
	assert.True(t, limiter.Allow("key2", ttl))
}
