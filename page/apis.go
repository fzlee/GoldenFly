package page

import (
	"github.com/gin-gonic/gin"
	"golden_fly/common"
	"net/http"
)

func PageList(c *gin.Context) {
	pagination := common.ParsePageAndSize(c)
	pages, _ := GetPages(&Page{}, &pagination)

	results := make([] *PageResponse, len(pages))
	for i := range pages {
		results[i] = (&PageSerializer{c, &pages[i]}).FullResponse()
	}

	c.JSON(http.StatusOK, gin.H{"data": results, "success": true})
}

func PagesPreview(c * gin.Context) {
	pagination := common.ParsePageAndSize(c)
	pages, _ := GetPages(&Page{}, &pagination)

	results := make([] *PageResponse, len(pages))
	for i := range pages {
		results[i] = (&PageSerializer{c, &pages[i]}).PreviewResponse()
	}

	c.JSON(http.StatusOK, gin.H{"data": results, "success": true})
}
