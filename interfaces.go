package main

import "time"

// NotificationService is an interface defining methods for sending notifications.
type NotificationService interface {
	SendNotification(userID string, message string, notificationType string) error
}

// RateLimiter is an interface defining methods for rate limiting.
type RateLimiter interface {
	Allow(key string, ttl time.Duration) bool
}

// Gateway is an interface defining methods for sending notifications.
type Gateway interface {
	Send(*Payload) error
}
