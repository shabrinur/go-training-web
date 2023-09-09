package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ShowPracticeIndex(c *gin.Context) {
	info := gin.H{
		"title": "Index",
		"year":  time.Now().Year(),
	}

	c.HTML(http.StatusOK, "index", info)
}

func ShowAbout(c *gin.Context) {
	info := gin.H{
		"title": "About",
		"year":  time.Now().Year(),
	}

	c.HTML(http.StatusOK, "about", info)
}

func ShowContact(c *gin.Context) {
	info := gin.H{
		"title": "Contact",
		"year":  time.Now().Year(),
	}

	c.HTML(http.StatusOK, "contact", info)
}
