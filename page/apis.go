package page

import (
	"github.com/gin-gonic/gin"
	"golden_fly/common"
	"net/http"
)

func ListPages (c *gin.Context) {
	pagination := common.ParsePageAndSize(c)
	pages, _ := GetPages(&Page{}, &pagination)

	results := make([] *PageResponse, len(pages))
	for i := range pages {
		results[i] = (&PageSerializer{c, &pages[i]}).FullResponse()
	}

	c.JSON(http.StatusOK, gin.H{"data": results, "success": true})
}

func RetrievePageMeta (c *gin.Context) {
	url :=  c.Param("name")
	page, err := GetPage(&Page{URL: url})
	if err != nil {
		common.AbortWithCode(c, http.StatusNotFound, common.CodeNotFound)
	}
	result := (&PageSerializer{c, &page}).MetaResponse()
	c.JSON(http.StatusOK, gin.H{"data": result, "success": true})

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


func PageComments (c *gin.Context) {
	pagination := common.ParsePageAndSize(c)
	url:= c.Param("url")
	page, err := GetPage(&Page{URL: url})

	if err != nil {
		common.ResponseWithCode(c, common.CodeNotFound)
	}
	comments, _ := GetComments(&Comment{PageID: page.ID}, &pagination, "id")

	results := make([] *CommentResponse, len(comments))
	for i := range comments {
		results[i] = (&CommentSerializer{c, &comments[i]}).CommentResponse(false)
	}

	c.JSON(http.StatusOK, gin.H{"data": results, "success": true})
}

func PageSideBar (c *gin.Context) {
	result := GenerateSideBar(c)
	c.JSON(http.StatusOK, gin.H{"data": result, "success": true})
}
