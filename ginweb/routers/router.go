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

	router.GET("/form/post", controllers.ShowFormA)
	router.POST("/form/post/result", controllers.ShowResultPost)

	router.GET("/form/query", controllers.ShowFormB)
	router.GET("/form/query/result", controllers.ShowResultQuery)

	router.GET("/form/path", controllers.ShowFormC)
	router.GET("/form/path/result/:input1/:input2", controllers.ShowResultPath)

	router.GET("/contact", controllers.ShowContact)

	return router
}
