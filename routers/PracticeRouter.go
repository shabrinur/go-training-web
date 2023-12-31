package routers

import (
	"idp-go/app/practice/controllers"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func SetupPracticeRouter(router *gin.Engine, render multitemplate.Renderer) (*gin.Engine, multitemplate.Renderer) {

	// set routing for main pages
	router.GET("/practice", controllers.ShowAbout)
	router.GET("/practice/about", controllers.ShowAbout)
	router.GET("/practice/contact", controllers.ShowContact)

	// call routing for sample forms
	router = setRoutingForSampleForms(router)

	router = setRoutingForToDoPages(router)

	// call practice pages render
	render = renderPracticeTemplates(render)

	return router, render
}

func setRoutingForSampleForms(r *gin.Engine) *gin.Engine {
	r.GET("/practice/form/post", controllers.ShowFormA)
	r.POST("/practice/form/post/result", controllers.ShowResultPost)

	r.GET("/practice/form/query", controllers.ShowFormB)
	r.GET("/practice/form/query/result", controllers.ShowResultQuery)

	r.GET("/practice/form/path", controllers.ShowFormC)
	r.GET("/practice/form/path/result/:input1/:input2", controllers.ShowResultPath)

	return r
}

func setRoutingForToDoPages(r *gin.Engine) *gin.Engine {
	r.GET("/practice/todo", controllers.ShowTodoList)

	r.GET("/practice/todo/create", controllers.ShowTodoCreate)
	r.POST("/practice/todo/create", controllers.PostTodoCreate)

	r.GET("/practice/todo/:todoId", controllers.ViewTodoDetail)

	r.GET("/practice/todo/:todoId/edit", controllers.ShowTodoUpdate)
	r.POST("/practice/todo/update", controllers.PostTodoUpdate)

	r.POST("/practice/todo/delete", controllers.PostTodoDelete)

	return r
}

func renderPracticeTemplates(render multitemplate.Renderer) multitemplate.Renderer {

	layoutPath := "app/practice/templates/layout.tmpl"
	viewPath := "app/practice/views"

	render.AddFromFiles("about", layoutPath, viewPath+"/about.html")

	render.AddFromFiles("post", layoutPath, viewPath+"/form/form-post.html")
	render.AddFromFiles("result-post", layoutPath, viewPath+"/form/result.html")
	render.AddFromFiles("query", layoutPath, viewPath+"/form/form-query.html")
	render.AddFromFiles("result-query", layoutPath, viewPath+"/form/result.html")
	render.AddFromFiles("path", layoutPath, viewPath+"/form/form-path.html")
	render.AddFromFiles("result-path", layoutPath, viewPath+"/form/result.html")

	render.AddFromFiles("todo-list", layoutPath, viewPath+"/todo/todo-list.html")
	render.AddFromFiles("todo-create", layoutPath, viewPath+"/todo/todo-create.html")
	render.AddFromFiles("todo-detail", layoutPath, viewPath+"/todo/todo-detail.html")
	render.AddFromFiles("todo-update", layoutPath, viewPath+"/todo/todo-update.html")
	render.AddFromFiles("todo-error", layoutPath, viewPath+"/todo/todo-error.html")

	render.AddFromFiles("contact", layoutPath, viewPath+"/contact.html")
	return render
}
