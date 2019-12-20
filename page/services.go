package page

import (
	"crypto/tls"
	"github.com/gin-gonic/gin"
	"golden_fly/common"
	"golden_fly/config"
	"gopkg.in/gomail.v2"
	"gopkg.in/russross/blackfriday.v2"
	"mime/multipart"
	"regexp"
	"strings"
	"time"
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
		"content": page.ContentDigest,
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

func TransactionDeletePage (page *Page) error {
	var err error
	// delete page, tags, comments
	tx := common.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	err = tx.Where("page_id = ?", page.ID).Delete(Comment{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Where("page_id = ?", page.ID).Delete(Tag{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Delete(&page).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}


func UpdateOrCreatePage(v *SavePageValidator) (*Page, error){
	var err error
	var page *Page
	if v.PageID != nil {
		page2, _ := GetPage(&Page{ID: *v.PageID})
		page = &page2
	} else {
		page = &Page{}
	}

	page.URL = v.URL
	page.Title = v.Title
	page.Content = v.Content
	page.MetaContent = v.MetaContent
	page.NeedKey = *v.NeedKey
	page.Tags = v.Tags
	page.Editor = v.Editor
	page.AllowComment = *v.AllowComment
	page.AllowVisit = *v.AllowVisit
	page.IsOriginal = *v.IsOriginal

	if page.NeedKey {
		page.Password = v.Password
	} else {
		page.Password = ""
	}

	// handle content
	if page.Editor == "html" {
		page.HTML = page.Content
	} else {
		page.HTML = string(blackfriday.Run([]byte(page.Content)))
	}
	r, _ := regexp.Compile("<.*?>")
	page.ContentDigest = r.ReplaceAllString(page.HTML, "")
	utf8String := []rune(page.ContentDigest)
	page.ContentDigest = string(utf8String[:common.MinINT(len(utf8String), 200)])
	// handle time
	now := time.Now()
	page.UpdateTime = &now
	if page.CreateTime == nil {
		page.CreateTime = &now
	}
	err = common.DB.Save(page).Error
	if err != nil {
		return nil, err
	}

	// handle tags
	if ! strings.HasPrefix(page.Tags, ",") {
		page.Tags = "," + page.Tags
	}
	if ! strings.HasSuffix(page.Tags, ",") {
		page.Tags = page.Tags + ","
	}

	err = common.DB.Save(page).Error
	if err != nil {
		return nil, err
	}
	return page, TransactionUpdatePageTags(page)
}


func TransactionUpdatePageTags (page *Page) error {
	common.DB.Begin()
	var err error
	// delete tags and create new
	tx := common.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	err = tx.Where("page_id = ?", page.ID).Delete(&Tag{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tags := strings.Split(page.Tags, ",")

	for i := range(tags) {
		tag := strings.Trim(tags[i], " ")
		if tag != "" {
			err = tx.Create(&Tag{
				Name:   tag,
				PageID: page.ID,
			}).Error

			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}


func CreateLink (v *CreateLinkValidator) error {
	link := &Link{
		Name:        v.Name,
		Href:        v.Href,
		Description: v.Description,
		CreateTime:  time.Now(),
		Display:     false,
	}

	return common.DB.Create(link).Error
}


func UpdateLink (link *Link, v *UpdateLinkValidator) {
	link.Display = *v.Display
	common.DB.Save(link)
}


func CreateMedia (header *multipart.FileHeader) (*Media, error){
	media := &Media{
		FileID: common.RandomString(20),
		FileName: header.Filename,
		Size: int(header.Size),
		Display: true,
		CreateTime: time.Now(),
		ContentType: header.Header["Content-Type"][0],
	}

	oldMedia, err := GetMedia(&Media{FileName: media.FileName}, "id desc")
	if err != nil {
		media.Version = 1
	} else {
		media.Version = oldMedia.Version + 1
		media.ID = oldMedia.ID
	}

	err = common.DB.Save(media).Error
	return media, err
}


func SendEmail (to string, subject string, content string) {
	conf := config.Get()
	if !conf.SendEmailReply {
		return
	}

	d := gomail.NewDialer(
		conf.EmailSMTPHost,
		conf.EmailSMTPPort,
		conf.EmailSMTPUsername,
		conf.EmailSMTPPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.SetHeader("From", conf.EmailSMTPUsername)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", content)

	// 异步执行邮件发送
	go d.DialAndSend(m)
}
