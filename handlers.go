package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	service NotificationService
}

// Handle handles a notification request.
func (nh *NotificationHandler) Handle(c *gin.Context) {
	var request NotificationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := nh.service.SendNotification(request.UserID, request.Message, request.Type); err != nil {
		status, body := handlerError(err)
		c.JSON(status, body)
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}

// handlerError handles errors from the notification service.
func handlerError(err error) (int, gin.H) {
	if errors.Is(err, ErrRateLimitExceeded) {
		return http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded"}
	}

	if errors.Is(err, ErrUnknownType) {
		return http.StatusBadRequest, gin.H{"error": "unknown notification type"}
	}

	if errors.Is(err, ErrGatewayFailure) {
		return http.StatusInternalServerError, gin.H{"error": "gateway failure"}
	}

	return http.StatusInternalServerError, gin.H{"error": "internal server error"}
}

// NewNotificationHandler initializes a new NotificationHandler.
func NewNotificationHandler(service NotificationService) *NotificationHandler {
	return &NotificationHandler{service: service}
}

// NotificationRequest represents the data structure for a notification request.
type NotificationRequest struct {
	UserID  string `json:"user_id"`
	Message string `json:"message"`
	Type    string `json:"type"`
}
