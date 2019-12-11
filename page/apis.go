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
	url :=  c.Param("url")
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


func PagesSearch(c *gin.Context) {
	tagname := c.DefaultQuery("tagname", "")
	if tagname == "" {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": make([] string, 0)})
		return
	}

	pagination := common.ParsePageAndSize(c)
	pages, _:= SearchPages(tagname, &pagination)

	results := make([] *PageResponse, len(pages))
	for i := range(pages) {
		results[i] = (&PageSerializer{c, &pages[i]}).PreviewResponse()
	}

	c.JSON(http.StatusOK, gin.H{"data": results, "success": true})

}


func CreateCommentView(c *gin.Context) {
	url :=  c.Param("url")

	page, err := GetPage(&Page{URL: url})
	if err != nil {
		common.AbortWithCode(c, http.StatusNotFound, common.CodeNotFound)
		return
	}

	var v CommentValidator
	if err := c.BindJSON(&v); err != nil {
		common.ResponseWithPanic(c, err)
		return
	}

	var to string
	if v.CommentID != nil {
		comment, err := GetComment(&Comment{ID: *v.CommentID})
		if err == nil {
			to = comment.Nickname
		}
	}

	CreateComment(&v, &to, c.ClientIP(), page.ID)
	c.JSON(http.StatusOK, gin.H{"success": true})
}
