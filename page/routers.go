package page

import "github.com/gin-gonic/gin"

func RegisterRouter(router *gin.RouterGroup) {
	router.GET("/articles/", PageList)
	router.GET("/articles/preview", PagesPreview)
}
