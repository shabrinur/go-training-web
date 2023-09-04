package routers

import (
	"ginweb/app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Routing
	router.GET("/", controllers.ShowIndex)
	router.GET("/about", controllers.ShowAbout)

	router.GET("/post", controllers.ShowFormA)
	router.POST("/result", controllers.ShowResultPost)

	router.GET("/query", controllers.ShowFormB)
	router.GET("/result", controllers.ShowResultQuery)

	router.GET("/path", controllers.ShowFormC)
	router.GET("/result/:input1/:input2", controllers.ShowResultPath)

	router.GET("/contact", controllers.ShowContact)

	return router
}
