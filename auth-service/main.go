package main

import (
	"net/http"
	_ "time"

	"github.com/gin-gonic/gin"

)

func main() {
	r := gin.Default()

	r.POST("/login", func(c *gin.Context) {
		var req struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// NOTE: In real app validate credentials and generate JWT
		// For demo return a dummy token
		token := "secret" + req.Username
		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	srv := &http.Server{
		Addr:    ":8081",
		Handler: r,
	}

	// graceful shutdown omitted for brevity in starter
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}