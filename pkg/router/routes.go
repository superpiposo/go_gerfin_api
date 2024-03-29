package router

import (
	"github.com/gin-gonic/gin"
	"github.com/superpiposo/go_gerfin_api/pkg/handler"
)

func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.GET("/usuarios", handler.GetUsers)
		v1.GET("/usuarios/:id", handler.GetUser)
		v1.DELETE("/usuarios", handler.DeleteUser)
		v1.PATCH("/usuarios", handler.UpdateUser)
		v1.POST("/usuarios", handler.CreateUser)
	}
}
