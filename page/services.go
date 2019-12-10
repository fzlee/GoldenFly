package page

import (
	"github.com/gin-gonic/gin"
	"golden_fly/common"
	"golden_fly/config"
)

func GenerateSideBar(c *gin.Context) *gin.H{
	return &gin.H{
		"announcement": generateSidebarAnnouncement(),
		"links": generateSidebarLinks(c),
		"comments": generateSidebarComments(c),
		"tags": generateSidebarTags(),
	}
}

func generateSidebarAnnouncement() *gin.H{
	conf := config.Get()
	page, err := GetPage(&Page{URL: conf.BlogAnnouncementURL})
	if err != nil {
		return &gin.H{}
	}

	return &gin.H{
		"url": page.URL,
		"content": page.ContentDigest[:common.MinINT(len(page.ContentDigest), 200)],
	}
}

func generateSidebarLinks(c *gin.Context) ([] *LinkResponse) {
	conf := config.Get()
	links, _ := GetLinks(&Link{Display: true}, &common.Pagination{Page: 1, Size: conf.BlogLinkCount})
	results := make([] *LinkResponse, len(links));
	for i := range(links) {
		results[i] = (&LinkSerializer{c, &links[i]}).getSidebarResponse()
	}
	return results
}

func generateSidebarComments(c *gin.Context) ([] *CommentResponse){
	conf := config.Get()
	comments, _ := GetComments(&Comment{}, &common.Pagination{Page: 1, Size:conf.BlogCommentCount}, "id desc")
	results := make([]*CommentResponse, len(comments))
	for i:= range(results) {
		results[i] = (&CommentSerializer{c, &comments[i]}).SidebarCommentResponse()
	}
	return results
}

func generateSidebarTags ()[]string {
	return GetDistinctTags()
}
