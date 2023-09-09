package main

import (
	"idp-go/routers"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	render := multitemplate.NewRenderer()

	// setup routing & render for practice pages
	router, render = routers.SetupPracticeRouter(router, render)

	// load asset
	router.Static("/assets", "./assets")

	// load all views
	router.HTMLRender = render

	router.Run(":8080")
}
