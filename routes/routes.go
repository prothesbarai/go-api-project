package routes

import (
	"go-api-project/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/products",controllers.GetProducts)

	// >>> For Igoner Fav Icons
	route.GET("/favicon.ico",func(ctx *gin.Context) {ctx.Status(204)})
}
