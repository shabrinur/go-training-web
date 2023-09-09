package routers

import (
	"idp-go/app/practice/controllers"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func SetupPracticeRouter(router *gin.Engine, render multitemplate.Renderer) (*gin.Engine, multitemplate.Renderer) {

	// set routing for main pages
	router.GET("/", controllers.ShowPracticeIndex)
	router.GET("/practice", controllers.ShowPracticeIndex)
	router.GET("/practice/contact", controllers.ShowContact)
	router.GET("/practice/about", controllers.ShowAbout)

	// call routing for sample forms
	router = setRoutingForSampleForms(router)

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

func renderPracticeTemplates(render multitemplate.Renderer) multitemplate.Renderer {

	layoutPath := "app/practice/templates/layout.tmpl"
	viewPath := "app/practice/views"

	render.AddFromFiles("index", layoutPath, viewPath+"/index.html")
	render.AddFromFiles("about", layoutPath, viewPath+"/about.html")

	render.AddFromFiles("post", layoutPath, viewPath+"/form/form-post.html")
	render.AddFromFiles("result-post", layoutPath, viewPath+"/form/result.html")

	render.AddFromFiles("query", layoutPath, viewPath+"/form/form-query.html")
	render.AddFromFiles("result-query", layoutPath, viewPath+"/form/result.html")

	render.AddFromFiles("path", layoutPath, viewPath+"/form/form-path.html")
	render.AddFromFiles("result-path", layoutPath, viewPath+"/form/result.html")

	render.AddFromFiles("contact", layoutPath, viewPath+"/contact.html")
	return render
}
