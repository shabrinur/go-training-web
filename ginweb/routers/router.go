package routers

import (
	"ginweb/app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Routing
	router.GET("/", controllers.Index)
	router.POST("/result", controllers.ResultValueOnBody)
	router.GET("/result", controllers.ResultValueOnQueryParam)
	router.GET("/result/:input1/:input2", controllers.ResultValueOnPathParam)

	return router
}
