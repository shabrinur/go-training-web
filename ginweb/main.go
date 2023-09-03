package main

import (
	"ginweb/routers"
)

func main() {

	// router init
	router := routers.SetupRouter()

	// load asset
	router.Static("/assets", "./assets")

	// load html
	router.LoadHTMLGlob("./app/views/*")

	router.Run(":8080")
}
