package controllers

import (
	"idp-go/app/practice/helpers"
	"idp-go/app/practice/models"
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

func ShowTodoCreate(c *gin.Context) {
	content := gin.H{
		"title": "To Do Create",
		"year":  time.Now().Year(),
	}
	c.HTML(http.StatusOK, "todo-create", content)
}

func PostTodoCreate(c *gin.Context) {
	title := c.PostForm("title")
	description := c.PostForm("description")

	todo := models.ToDo{
		Title:       title,
		Description: description,
	}

	err := helpers.Save(todo)
	if err != nil {
		errInfo := gin.H{
			"title":   "Error",
			"year":    time.Now().Year(),
			"process": "SAVE_TODO_ENTRY_TO_DB",
			"message": err.Error(),
		}
		c.HTML(http.StatusInternalServerError, "todo-error", errInfo)
		return
	}
	c.Redirect(http.StatusSeeOther, "/practice/todo")
}
