package user

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.RouterGroup) {
	router.GET("/users/me/", UsersMeView)
	router.POST("/login/", LoginView)

	adminGroup := router.Group("/")
	adminGroup.Use(AdminRequired)
	adminGroup.GET("/users/", UsersListView)
}
