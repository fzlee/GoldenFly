package other

import (
	"github.com/gin-gonic/gin"
	"golden_fly/config"
)


func RegisterRouter(engine *gin.Engine) {
	conf := config.Get()
	// sitemap
	engine.GET("/sitemap.xml", GenerateSitemapView)
	// rss
	engine.GET("/rss", GenerateRSSView)
	// static folder
	engine.StaticFS("/media", gin.Dir(conf.MediaFolder, false))
}
