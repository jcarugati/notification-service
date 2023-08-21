package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	handler := wire()

	routes(r, handler)

	if err := r.Run(); err != nil {
		panic(err)
	}
}

// routes registers routes for the application.
func wire() *NotificationHandler {
	manifest, err := LoadFile("./manifest.yaml")
	if err != nil {
		panic(err)
	}

	gateway := NewLocalGateway()
	limiter := NewLocalRateLimiter()

	service := NewNotificationService(limiter, gateway, manifest)

	handler := NewNotificationHandler(service)

	return handler
}
