package routers

import (
	"idp-go/app/task/controllers"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func SetupTaskRouter(router *gin.Engine, render multitemplate.Renderer) (*gin.Engine, multitemplate.Renderer) {

	// set routing for main pages
	router.GET("/", controllers.ShowIndex)
	router.POST("/welcome", controllers.ShowWelcome)

	// set handling for no route & no method
	router = setRoutingErrors(router)

	// call practice pages render
	render = renderTaskTemplates(render)

	return router, render
}

func setRoutingErrors(r *gin.Engine) *gin.Engine {
	r.NoRoute(controllers.ShowNotFound)
	r.NoMethod(controllers.ShowMethodNotAllowed)

	return r
}

func renderTaskTemplates(render multitemplate.Renderer) multitemplate.Renderer {
	layoutPath := "app/task/templates/layout.tmpl"
	viewPath := "app/task/views"

	render.AddFromFiles("task-index", layoutPath, viewPath+"/index.html")
	render.AddFromFiles("welcome", layoutPath, viewPath+"/welcome.html")
	render.AddFromFiles("not-found", layoutPath, viewPath+"/error.html")
	render.AddFromFiles("not-allowed", layoutPath, viewPath+"/error.html")

	return render
}
