package user

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.RouterGroup) {
	router.GET("/users/me/", UsersMe)
	router.POST("/login/", Login)

	adminGroup := router.Group("/")
	adminGroup.Use(AdminRequired)
	adminGroup.GET("/users/", UsersList)
}
