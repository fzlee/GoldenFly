package page

import (
	"github.com/gin-gonic/gin"
	"golden_fly/user"
)

func RegisterRouter(router *gin.RouterGroup) {
	router.GET("/articles/", ListPages)
	router.GET("/articles-preview/", PagesPreview)
	router.GET("/articles-sidebar/", PageSideBar)
	router.GET("/articles-search/", PagesSearch)
	router.POST("/articles-inplace/", PagesInPlace)
	router.GET("/articles/:url/meta/", RetrievePageMeta)
	router.GET("/articles/:url/comments/", PageComments)
	router.POST("/articles/:url/comments/", CreateCommentView)

	adminGroup := router.Group("/")
	adminGroup.Use(user.AdminRequired)
	adminGroup.GET("/articles/:url/", RetrievePage)
	adminGroup.DELETE("/articles/:url/", DeletePage)
	adminGroup.PUT("/articles/save/", SavePage)
	adminGroup.GET("/comments/", ListComments)
	adminGroup.DELETE("/comments/:id/", DeleteComment)
	adminGroup.GET("/links/", ListLinksView)
	adminGroup.POST("/links/", CreateLinkView)
	adminGroup.DELETE("/links/:id/", DeleteLinkView)
	adminGroup.PUT("/links/:id/", UpdateLinkView)
	adminGroup.GET("/medias/", ListMediasView)
	adminGroup.DELETE("/medias/:id/", DeleteMediaView)
}
