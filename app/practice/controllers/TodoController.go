package controllers

import (
	"idp-go/app/practice/helpers"
	"idp-go/app/practice/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func ShowTodoList(c *gin.Context) {
	todos, err := helpers.GetList()
	if err != nil {
		errorHandling("GET_TODO_LIST_FROM_DB", err, c)
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
		errorHandling("SAVE_TODO_ENTRY_TO_DB", err, c)
		return
	}
	c.Redirect(http.StatusSeeOther, "/practice/todo")
}

func ViewTodoDetail(c *gin.Context) {
	id := c.Param("todoId")
	val, err := strconv.Atoi(id)
	if err != nil {
		errorHandling("PARSE_TODO_ID", err, c)
		return
	}

	todo, err := helpers.Get(val)
	if err != nil {
		errorHandling("GET_TODO_ENTRY_FROM_DB", err, c)
		return
	}

	content := gin.H{
		"title": "To Do Detail",
		"year":  time.Now().Year(),
		"todo":  todo,
	}
	c.HTML(http.StatusOK, "todo-detail", content)
}

func ShowTodoUpdate(c *gin.Context) {
	id := c.Param("todoId")
	val, err := strconv.Atoi(id)
	if err != nil {
		errorHandling("PARSE_TODO_ID", err, c)
		return
	}

	todo, err := helpers.Get(val)
	if err != nil {
		errorHandling("GET_TODO_ENTRY_FROM_DB", err, c)
		return
	}

	content := gin.H{
		"title": "To Do Update",
		"year":  time.Now().Year(),
		"todo":  todo,
	}
	c.HTML(http.StatusOK, "todo-update", content)
}

func PostTodoUpdate(c *gin.Context) {
	id := c.PostForm("todoId")
	title := c.PostForm("title")
	description := c.PostForm("description")

	val, err := strconv.Atoi(id)
	if err != nil {
		errorHandling("PARSE_TODO_ID", err, c)
		return
	}

	todo := models.ToDo{
		Title:       title,
		Description: description,
	}

	err = helpers.Update(val, todo)
	if err != nil {
		errorHandling("UPDATE_TODO_ENTRY_TO_DB", err, c)
		return
	}
	c.Redirect(http.StatusSeeOther, "/practice/todo")
}

func errorHandling(process string, err error, c *gin.Context) {
	errInfo := gin.H{
		"title":   "Error",
		"year":    time.Now().Year(),
		"process": process,
		"message": err.Error(),
	}
	c.HTML(http.StatusInternalServerError, "todo-error", errInfo)
}
