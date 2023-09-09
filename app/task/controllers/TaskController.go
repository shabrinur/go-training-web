package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ShowIndex(c *gin.Context) {
	info := gin.H{
		"title": "Index",
		"year":  time.Now().Year(),
	}

	c.HTML(http.StatusOK, "task-index", info)
}

func ShowWelcome(c *gin.Context) {
	inputName := c.PostForm("inputName")

	info := gin.H{
		"title":     "Welcome",
		"inputName": inputName,
		"year":      time.Now().Year(),
	}

	c.HTML(http.StatusOK, "welcome", info)
}

func ShowNotFound(c *gin.Context) {
	info := gin.H{
		"title":   "404 Not Found",
		"errName": "Page Not Found",
		"errMsg":  "Sorry, we can't find the page you're looking for",
		"year":    time.Now().Year(),
	}
	c.HTML(http.StatusNotFound, "not-found", info)
}

func ShowMethodNotAllowed(c *gin.Context) {
	info := gin.H{
		"title":   "405 Method Not Allowed",
		"errName": "Method Not Allowed",
		"errMsg":  "Sorry, you can't access page with this method",
		"year":    time.Now().Year(),
	}
	c.HTML(http.StatusMethodNotAllowed, "not-allowed", info)
}
