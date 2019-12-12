package page

import (
	"github.com/gin-gonic/gin"
	"golden_fly/common"
	"net/http"
	"strconv"
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


func DeletePage (c *gin.Context) {
	var err error
	var page Page
	url := c.Param("url")
	page, err = GetPage(&Page{URL: url})
	if err != nil {
		common.AbortWithCode(c, http.StatusNotFound, common.CodeNotFound)
	}
	err = TransactionDeletePage(&page)
	if err != nil {
		common.ResponseWithPanic(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
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

func RetrievePage(c *gin.Context) {
	url := c.Param("url")

	page, err := GetPage(&Page{URL: url})
	if err != nil {
		common.AbortWithCode(c, http.StatusNotFound, common.CodeNotFound)
	}
	result := (&PageSerializer{c, &page}).FullResponse()
	c.JSON(http.StatusOK, gin.H{"data": result, "success": true})
}

func PagesInPlace(c *gin.Context) {
	var v InPlaceValidator
	if err := c.BindJSON(&v); err != nil {
		common.ResponseWithValidation(c, err)
		return
	}

	count := 1
	common.DB.Model(&Page{}).Where(&Page{URL:v.URL}).Count(&count)

	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{
		"in_place": count > 0,
	}})
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


func SavePage (c *gin.Context) {

	var err error
	var v SavePageValidator
	if err = c.BindJSON(&v); err != nil {
		common.ResponseWithValidation(c, err)
		return
	}
	var page Page
	if err = UpdateOrCreatePage(&v, &page); err != nil {
		common.ResponseWithValidation(c, err)
		return
	}

	s := PageSerializer{c, &page}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": s.FullResponse(),
	})

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
		common.ResponseWithValidation(c, err)
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


func ListComments(c *gin.Context) {
	pagination := common.ParsePageAndSize(c)
	comments, _ := GetComments(&Comment{}, &pagination, "id desc")

	results := make([] *CommentResponse, len(comments))
	for i := range comments {
		results[i] = (&CommentSerializer{c, &comments[i]}).CommentResponse(true)
	}

	c.JSON(http.StatusOK, gin.H{"data": results, "success": true})
}


func DeleteComment (c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	comment, _ := GetComment(&Comment{ID: id})

	common.DB.Delete(&comment)
	c.JSON(http.StatusOK, gin.H{"success": true})
}


func ListLinksView(c *gin.Context) {
	pagination := common.ParsePageAndSize(c)
	links , _ := GetLinks(&Link{}, &pagination)

	results := make([] *LinkResponse, len(links))
	for i := range links{
		results[i] = (&LinkSerializer{c, &links[i]}).FullResponse()
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": results})
}


func CreateLinkView(c *gin.Context) {
	var v CreateLinkValidator
	var err error
	if err = c.BindJSON(&v); err != nil {
		common.ResponseWithValidation(c, err)
		return
	}

	err = CreateLink(&v)
	if err != nil {
		common.ResponseWithPanic(c, err)
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

func DeleteLinkView (c *gin.Context) {
	var err error
	var id int
	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		common.ResponseWithCode(c, common.CodeNotFound)
		return
	}

	var link Link
	link, err = GetLink(&Link{ID: id})
	if err != nil {
		common.ResponseWithCode(c, common.CodeNotFound)
		return
	}

	err = common.DB.Delete(&link).Error
	if err != nil {
		common.ResponseWithPanic(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}


func UpdateLinkView (c *gin.Context) {
	var err error
	var id int
	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		common.ResponseWithCode(c, common.CodeNotFound)
		return
	}

	var link Link
	link, err = GetLink(&Link{ID: id})
	if err != nil {
		common.ResponseWithCode(c, common.CodeNotFound)
		return
	}

	var v UpdateLinkValidator
	if err = c.BindJSON(&v); err !=nil {
		common.ResponseWithPanic(c, err)
	}

	UpdateLink(&link, &v)
	c.JSON(http.StatusOK, gin.H{"success": true})

}


func ListMediasView(c *gin.Context) {
	pagination := common.ParsePageAndSize(c)
	medias , _ := GetMedias(&Media{}, &pagination)

	results := make([] *MediaResponse, len(medias))
	for i := range medias{
		results[i] = (&MediaSerializer{c, &medias[i]}).FullResponse()
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": results})
}


func DeleteMediaView (c *gin.Context) {
	var id int
	var err error
	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		common.ResponseWithCode(c, common.CodeNotFound)
		return
	}

	var media Media
	media, err = GetMedia(&Media{ID: id})
	if err != nil {
		common.ResponseWithCode(c, common.CodeNotFound)
		return
	}

	media.DeleteLocalFile()
	err = common.DB.Delete(&media).Error
	c.JSON(http.StatusOK, gin.H{"success": true})
}
