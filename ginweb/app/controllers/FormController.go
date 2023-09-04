package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ShowFormA(c *gin.Context) {
	info := gin.H{
		"title": "Form A",
		"year":  time.Now().Year(),
	}

	c.HTML(http.StatusOK, "post", info)
}

func ShowFormB(c *gin.Context) {
	info := gin.H{
		"title": "Form B",
		"year":  time.Now().Year(),
	}

	c.HTML(http.StatusOK, "query", info)
}

func ShowFormC(c *gin.Context) {
	info := gin.H{
		"title": "Form C",
		"year":  time.Now().Year(),
	}

	c.HTML(http.StatusOK, "path", info)
}

func ShowResultPost(c *gin.Context) {
	input1 := c.PostForm("input1A")
	input2 := c.PostForm("input2A")

	result := gin.H{
		"title":    "Form A Result",
		"formName": "Form A",
		"input1":   input1,
		"input2":   input2,
		"year":     time.Now().Year(),
	}

	c.HTML(http.StatusOK, "result-post", result)
}

func ShowResultQuery(c *gin.Context) {
	input1 := c.Query("input1B")
	input2 := c.Query("input2B")

	result := gin.H{
		"title":    "Form B Result",
		"formName": "Form B",
		"input1":   input1,
		"input2":   input2,
		"year":     time.Now().Year(),
	}

	c.HTML(http.StatusOK, "result-query", result)
}

func ShowResultPath(c *gin.Context) {
	input1 := c.Param("input1")
	input2 := c.Param("input2")

	result := gin.H{
		"title":    "Form C Result",
		"formName": "Form C",
		"input1":   input1,
		"input2":   input2,
		"year":     time.Now().Year(),
	}

	c.HTML(http.StatusOK, "result-path", result)
}
