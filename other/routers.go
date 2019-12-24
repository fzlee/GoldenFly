package other

import (
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"golden_fly/config"
	"time"
)

func RegisterRouter(engine *gin.Engine, store *persistence.InMemoryStore) {
	conf := config.Get()
	// sitemap
	engine.GET("/sitemap.xml", cache.CachePage(store, time.Minute*5, GenerateSitemapView))
	// cache rss
	engine.GET("/rss", cache.CachePage(store, time.Minute*5, GenerateRSSView))
	// static folder
	engine.StaticFS("/media", gin.Dir(conf.MediaFolder, false))
}
