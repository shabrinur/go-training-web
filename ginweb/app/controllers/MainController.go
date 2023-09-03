package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	info := gin.H{
		"title":       "Gin Web Apps",
		"description": "Hello World! :)",
	}

	c.HTML(http.StatusOK, "index.html", info)
}

func ResultValueOnBody(c *gin.Context) {
	input1 := c.PostForm("input1A")
	input2 := c.PostForm("input2A")

	result := gin.H{
		"title":  "Gin Web Apps - Show Result Value On Body",
		"input1": input1,
		"input2": input2,
	}

	c.HTML(http.StatusOK, "result.html", result)
}

func ResultValueOnQueryParam(c *gin.Context) {
	input1 := c.Query("input1B")
	input2 := c.Query("input2B")

	result := gin.H{
		"title":  "Gin Web Apps - Show Result Value On Query Param",
		"input1": input1,
		"input2": input2,
	}

	c.HTML(http.StatusOK, "result.html", result)
}

func ResultValueOnPathParam(c *gin.Context) {
	input1 := c.Param("input1")
	input2 := c.Param("input2")

	result := gin.H{
		"title":  "Gin Web Apps - Show Result Value On Path Param",
		"input1": input1,
		"input2": input2,
	}

	c.HTML(http.StatusOK, "result.html", result)
}
