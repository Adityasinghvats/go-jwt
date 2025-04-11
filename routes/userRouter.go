package routes

import (
	controller "github.com/adix/gojwt/controllers"
	"github.com/adix/gojwt/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middlewares.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users:id", controller.GetUser())
}
