package page

import (
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"golden_fly/user"
	"time"
)

func RegisterRouter(router *gin.RouterGroup, store *persistence.InMemoryStore) {
	router.GET("/articles-preview/", PagesPreview)
	router.GET("/articles-search/", PagesSearchView)
	router.POST("/articles-inplace/", PagesInPlaceView)
	router.GET("/articles/:url/meta/", RetrievePageMetaView)
	router.POST("/articles/:url/meta/", GetPageByPasswordView)
	router.GET("/articles/:url/comments/", PageCommentsView)
	router.POST("/articles/:url/comments/", CreateCommentView)

	// cache
	router.GET("/articles-sidebar/", cache.CachePage(store, time.Second * 10, PageSideBarView))

	adminGroup := router.Group("/")
	adminGroup.Use(user.AdminRequired)
	adminGroup.GET("/articles/", ListPagesView)
	adminGroup.GET("/articles/:url/", RetrievePageView)
	adminGroup.DELETE("/articles/:url/", DeletePageView)
	adminGroup.PUT("/articles/save/", SavePageView)
	adminGroup.GET("/comments/", ListCommentsView)
	adminGroup.DELETE("/comments/:id/", DeleteCommentView)
	adminGroup.GET("/links/", ListLinksView)
	adminGroup.POST("/links/", CreateLinkView)
	adminGroup.DELETE("/links/:id/", DeleteLinkView)
	adminGroup.PUT("/links/:id/", UpdateLinkView)
	adminGroup.GET("/medias/", ListMediasView)
	adminGroup.DELETE("/medias/:id/", DeleteMediaView)
	adminGroup.POST("/medias/upload/", UploadMediaView)
}
