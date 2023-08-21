package main

import "github.com/gin-gonic/gin"

func routes(r *gin.Engine, handler *NotificationHandler) {
	r.GET("/health", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/api/v1")
	{
		v1.POST("/notification", handler.Handle)
	}
}
