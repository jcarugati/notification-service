package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNotificationHandler_Handle(t *testing.T) {
	r := gin.Default()
	routes(r, makeHandler())

	t.Run("should send notification", func(t *testing.T) {
		w := httptest.NewRecorder()

		requestBody := &NotificationRequest{
			UserID:  "user1",
			Message: "Test Message",
			Type:    "test1",
		}

		b, _ := json.Marshal(requestBody)

		req, _ := http.NewRequest("POST", "/api/v1/notification", bytes.NewBuffer(b))
		r.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})

	t.Run("should return 400 when request body is invalid", func(t *testing.T) {
		w := httptest.NewRecorder()

		req, _ := http.NewRequest("POST", "/api/v1/notification", bytes.NewBuffer([]byte("invalid")))
		r.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	})

	t.Run("should return 400 when notification type is unknown", func(t *testing.T) {
		w := httptest.NewRecorder()

		requestBody := &NotificationRequest{
			UserID:  "user1",
			Message: "Test Message",
			Type:    "unknown",
		}

		b, _ := json.Marshal(requestBody)

		req, _ := http.NewRequest("POST", "/api/v1/notification", bytes.NewBuffer(b))
		r.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	})

	t.Run("should return 429 when rate limit is exceeded", func(t *testing.T) {
		var wa []*httptest.ResponseRecorder

		for range [2]int{} {
			w := httptest.NewRecorder()

			requestBody := &NotificationRequest{
				UserID:  "user2",
				Message: "Test Message",
				Type:    "test1",
			}

			b, _ := json.Marshal(requestBody)

			req, _ := http.NewRequest("POST", "/api/v1/notification", bytes.NewBuffer(b))
			r.ServeHTTP(w, req)

			wa = append(wa, w)
		}

		assert.Equal(t, 200, wa[0].Code)
		assert.Equal(t, 429, wa[1].Code)
	})
}

func makeHandler() *NotificationHandler {
	manifest := &Manifest{
		Version: 0,
		Rules: []Values{
			{
				Name:        "TestRule",
				Type:        "test1",
				TTL:         "1s",
				MaxAttempts: 1,
			},
		},
	}

	gateway := NewLocalGateway()
	limiter := NewLocalRateLimiter()

	service := NewNotificationService(limiter, gateway, manifest)

	handler := NewNotificationHandler(service)

	return handler
}
