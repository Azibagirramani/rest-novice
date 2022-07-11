package main

import (
	"net/http"
	"restapi/controllers"

	"github.com/gin-gonic/gin"
)

// no found handler
func noRoute(cxt *gin.Context) {
	cxt.JSON(http.StatusNotFound, map[string]string{
		"message": "No endpoint found",
	})
}

func main() {
	port := ":7777"

	server := gin.Default()
	server.Use(gin.Logger())

	api := server.Group("/api")

	{
		api.GET("/content", controllers.GetAll)
		api.POST("/create-content", controllers.CreateContent)
	}

	server.NoRoute(noRoute)

	server.Run(port)
}
