package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotificationService_SendNotification(t *testing.T) {

	t.Run("should send notification", func(t *testing.T) {
		service := makeService()
		err := service.SendNotification("user_id", "message", "test1")
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("should not send notification", func(t *testing.T) {
		service := makeService()
		err := service.SendNotification("user_id", "message", "test2")
		if err == nil {
			t.Error("expected error")
		}
		assert.ErrorIs(t, err, ErrUnknownType)
	})

	t.Run("should not send notification because of TTL", func(t *testing.T) {
		service := makeService()
		err := service.SendNotification("user_id", "message", "test1")
		if err != nil {
			t.Error(err)
		}

		err = service.SendNotification("user_id", "message", "test1")
		if err == nil {
			t.Error("expected error")
		}
		assert.ErrorIs(t, err, ErrRateLimitExceeded)
	})
}

func makeService() NotificationService {
	return NewNotificationService(
		NewLocalRateLimiter(),
		NewLocalGateway(),
		&Manifest{
			Version: 0,
			Rules: []Values{
				{
					Name: "TestRule",
					Type: "test1",
					TTL:  "1s",
				},
			},
		})
}
