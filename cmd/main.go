package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/ping", getting)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	
	http.ListenAndServe(":3006", router)
	// router.Run(":3000") for a hard coded port
}

func getting(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}