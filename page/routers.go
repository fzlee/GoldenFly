package page

import "github.com/gin-gonic/gin"

func RegisterRouter(router *gin.RouterGroup) {
	router.GET("/articles/", ListPages)
	router.GET("/articles-preview/", PagesPreview)
	router.GET("/articles-sidebar/", PageSideBar)
	router.GET("/articles-search/", PagesSearch)
	router.GET("/articles/:url/meta/", RetrievePageMeta)
	router.GET("/articles/:url/comments/", PageComments)
	router.POST("/articles/:url/comments/", CreateCommentView)
}
