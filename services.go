package main

import (
	"fmt"
	"time"
)

var (
	// ErrRateLimitExceeded is returned when a rate limit is exceeded.
	ErrRateLimitExceeded = fmt.Errorf("rate limit exceeded")
	// ErrGatewayFailure is returned when there's a failure in sending through the gateway.
	ErrGatewayFailure = fmt.Errorf("gateway failure")
	// ErrUnknownType is returned when an unknown notification type is provided.
	ErrUnknownType = fmt.Errorf("unknown notification type")
)

// Rule represents a rate limiting rule with a specified TTL.
type Rule struct {
	TTL         time.Duration
	MaxAttempts int
}

// Rules is a map that represents a collection of rules indexed by their type.
type Rules map[string]*Rule

// notificationService is an internal implementation of the NotificationService interface.
type notificationService struct {
	Limiter RateLimiter
	Gateway Gateway
	Rules   Rules
}

// SendNotification sends a notification to a user based on a given type.
// It checks rate limiting rules and sends through the gateway.
func (ns *notificationService) SendNotification(userID, message, notificationType string) error {
	// Get rate limiting rules for a given notification type.
	rule, err := ns.getRules(notificationType)
	if err != nil {
		return err
	}

	// Create a composite key from userID and notificationType.
	key := createKey(userID, notificationType)

	// Check if request is allowed based on the key and TTL.
	if !ns.Limiter.Allow(key, rule) {
		return ErrRateLimitExceeded
	}

	// Send notification through the gateway.
	payload := &Payload{UserID: userID, Message: message}
	if err := ns.Gateway.Send(payload); err != nil {
		return ErrGatewayFailure
	}

	return nil
}

// createKey creates a composite key from userID and notificationType.
func createKey(userID string, notificationType string) string {
	return fmt.Sprintf("%s:%s", userID, notificationType)
}

// getRules retrieves rate limiting rules for a given notification type.
func (ns *notificationService) getRules(notificationType string) (*Rule, error) {
	val, ok := ns.Rules[notificationType]
	if !ok {
		return nil, ErrUnknownType
	}
	return val, nil
}

// NewNotificationService initializes a new notification service with the provided limiter, gateway, and manifest.
func NewNotificationService(limiter RateLimiter, gateway Gateway, manifest *Manifest) NotificationService {
	rules := make(Rules)

	for _, rule := range manifest.Rules {
		ttl, err := time.ParseDuration(rule.TTL)
		if err != nil {
			panic(err)
		}

		rules[rule.Type] = &Rule{TTL: ttl, MaxAttempts: rule.MaxAttempts}
	}

	return &notificationService{Limiter: limiter, Gateway: gateway, Rules: rules}
}
