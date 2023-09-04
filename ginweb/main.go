package main

import (
	"ginweb/routers"

	"github.com/gin-contrib/multitemplate"
)

func main() {

	// router init
	router := routers.SetupRouter()

	// load asset
	router.Static("/assets", "./assets")

	// load html
	router.HTMLRender = renderTemplates()

	router.Run(":8080")
}

func renderTemplates() multitemplate.Renderer {
	render := multitemplate.NewRenderer()
	render.AddFromFiles("index", "app/templates/layout.tmpl", "app/views/index.html")
	render.AddFromFiles("about", "app/templates/layout.tmpl", "app/views/about.html")

	render.AddFromFiles("post", "app/templates/layout.tmpl", "app/views/form-post.html")
	render.AddFromFiles("result-post", "app/templates/layout.tmpl", "app/views/result.html")

	render.AddFromFiles("query", "app/templates/layout.tmpl", "app/views/form-query.html")
	render.AddFromFiles("result-query", "app/templates/layout.tmpl", "app/views/result.html")

	render.AddFromFiles("path", "app/templates/layout.tmpl", "app/views/form-path.html")
	render.AddFromFiles("result-path", "app/templates/layout.tmpl", "app/views/result.html")

	render.AddFromFiles("contact", "app/templates/layout.tmpl", "app/views/contact.html")
	return render
}
