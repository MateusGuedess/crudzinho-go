package routes

import (
	"github.com/MateusGuedess/crudzinho-go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.RouterGroup) {
	router.GET("/user/:userId", controller.FindUserByID)
	router.GET("/user", controller.FindUserByEmail)
	router.POST("/user", controller.CreateUser)
	router.PUT("/user/:userId", controller.UpdateUser)
	router.DELETE("/user/:userId", controller.DeleteUser)
}
