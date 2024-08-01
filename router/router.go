package router

import (
	"cervejarias-app/handlers"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()
	router.POST("/", handlers.SearchHandlerGin)
	router.StaticFile("/", "static/search_form.html")
	router.GET("/api", handlers.GetExternalApi)
	router.Run(":8080")
}
