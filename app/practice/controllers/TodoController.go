package controllers

import (
	"idp-go/app/practice/helpers"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ShowTodoList(c *gin.Context) {
	todos, err := helpers.GetList()
	if err != nil {
		errInfo := gin.H{
			"title":   "Error",
			"year":    time.Now().Year(),
			"process": "GET_TODO_LIST_FROM_DB",
			"message": err.Error(),
		}
		c.HTML(http.StatusInternalServerError, "todo-error", errInfo)
		return
	}
	content := gin.H{
		"title": "To Do List",
		"year":  time.Now().Year(),
		"todos": todos,
	}
	c.HTML(http.StatusOK, "todo-list", content)
}
