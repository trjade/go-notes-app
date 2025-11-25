package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Note struct {
	ID      string `json:"id"`
	OwnerID string `json:"owner_id"`
	Title   string `json:"title"`
	Body    string `json:"body"`
}

var notes = []Note{
	{ID: "1", OwnerID: "1", Title: "Buy milk", Body: "Remember to buy milk"},
}

func main() {
	r := gin.Default()

	r.GET("/notes", func(c *gin.Context) {
		c.JSON(http.StatusOK, notes)
	})

	r.POST("/notes", func(c *gin.Context) {
		var n Note
		if err := c.ShouldBindJSON(&n); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		n.ID = "new-id"
		notes = append(notes, n)
		c.JSON(http.StatusCreated, n)
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	if err := r.Run(":8083"); err != nil {
		panic(err)
	}
}