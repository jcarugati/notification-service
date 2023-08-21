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
	rule := &Rule{MaxAttempts: 1, TTL: time.Second}

	// First call to Allow for a given key should return true.
	assert.True(t, limiter.Allow("key1", rule))
}

func TestDisallowWithinTTL(t *testing.T) {
	limiter := NewLocalRateLimiter()
	rule := &Rule{MaxAttempts: 1, TTL: time.Second}

	// Set an item with TTL.
	limiter.Allow("key1", rule)

	// Within TTL, Allow should return false.
	assert.False(t, limiter.Allow("key1", rule))
}

func TestAllowAfterTTLExpires(t *testing.T) {
	limiter := NewLocalRateLimiter()
	rule := &Rule{MaxAttempts: 1, TTL: 100 * time.Millisecond}

	// Set an item with TTL.
	limiter.Allow("key1", rule)

	// Wait for the TTL to expire.
	time.Sleep(rule.TTL + 50*time.Millisecond)

	// After TTL expires, Allow should return true.
	assert.True(t, limiter.Allow("key1", rule))
}

func TestAllowMaxRetries(t *testing.T) {
	var boolArr []bool

	limiter := NewLocalRateLimiter()
	rule := &Rule{MaxAttempts: 4, TTL: time.Second}

	// Set an item with TTL.
	for i := 0; i < rule.MaxAttempts+1; i++ {
		boolArr = append(boolArr, limiter.Allow("key1", rule))
	}

	// Allow should return true for the first 4 calls.
	assert.True(t, boolArr[0])
	assert.True(t, boolArr[1])
	assert.True(t, boolArr[2])
	assert.True(t, boolArr[3])
	assert.False(t, boolArr[4])
}

func TestAllowMaxRetriesAndTTL(t *testing.T) {
	var boolArr []bool

	limiter := NewLocalRateLimiter()
	rule := &Rule{MaxAttempts: 2, TTL: time.Second}

	for i := 0; i < rule.MaxAttempts+1; i++ {
		boolArr = append(boolArr, limiter.Allow("key1", rule))
	}

	time.Sleep(time.Second)

	for i := 0; i < rule.MaxAttempts+1; i++ {
		boolArr = append(boolArr, limiter.Allow("key1", rule))
	}

	// Allow should return true for the first 4 calls.
	assert.True(t, boolArr[0])
	assert.True(t, boolArr[1])
	assert.False(t, boolArr[2])
	assert.True(t, boolArr[3])
	assert.True(t, boolArr[4])
	assert.False(t, boolArr[5])
}

func TestDifferentKeys(t *testing.T) {
	limiter := NewLocalRateLimiter()
	ttl := time.Second
	rule := &Rule{MaxAttempts: 1, TTL: ttl}

	// Set an item with TTL for key1.
	limiter.Allow("key1", rule)

	// Different key should not be affected by the TTL of key1.
	assert.True(t, limiter.Allow("key2", rule))
}
