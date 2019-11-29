package user

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.RouterGroup) {
	router.GET("/users/me/", UsersMe)
}
