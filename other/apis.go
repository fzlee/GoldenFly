package other

import (
	"github.com/gin-gonic/gin"
	"golden_fly/common"
	"golden_fly/page"
	"net/http"
)

func GenerateSitemapView(c *gin.Context) {

	pages, _ := page.GetPages(&page.Page{AllowVisit: true, NeedKey: false}, &common.Pagination{Page: 1, Size: 100000})
	stm := GenerateSitemap(&pages)
	c.Data(http.StatusOK, "application/xml", []byte(stm.XMLContent()))

}

func GenerateRSSView(c *gin.Context) {
	pages, _ := page.GetPages(&page.Page{AllowVisit: true}, &common.Pagination{Page: 1, Size: 100})
	content, _ := GenerateRSS(&pages)

	c.Data(http.StatusOK, "application/xml", []byte(content))
}
