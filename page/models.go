package page

import (
	"database/sql"
	"golden_fly/common"
	"time"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

type Page struct {
	ID            int       `gorm:"column:id;primary_key" json:"id"`
	URL           string    `gorm:"column:url" json:"url"`
	Title         string    `gorm:"column:title" json:"title"`
	ContentDigest string    `gorm:"column:content_digest" json:"content_digest"`
	Content       string    `gorm:"column:content" json:"content"`
	Keywords      string    `gorm:"column:keywords" json:"keywords"`
	MetaContent   string    `gorm:"column:metacontent" json:"metacontent"`
	CreateTime    time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime    time.Time `gorm:"column:update_time" json:"update_time"`
	NeedKey       bool       `gorm:"column:need_key" json:"need_key"`
	Password      string    `gorm:"column:password" json:"password"`
	Tags          string    `gorm:"column:tags" json:"tags"`
	Editor        string    `gorm:"column:editor" json:"editor"`
	AllowVisit    bool       `gorm:"column:allow_visit" json:"allow_visit"`
	AllowComment  bool       `gorm:"column:allow_comment" json:"allow_comment"`
	IsOriginal    bool       `gorm:"column:is_original" json:"is_original"`
	NumLookup     int       `gorm:"column:num_lookup" json:"num_lookup"`
	HTML          string    `gorm:"column:html" json:"html"`
}

// TableName sets the insert table name for this struct type
func (p *Page) TableName() string {
	return "page"
}


func GetPage (condition interface{}) (Page, error){
	var page Page;
	err := common.DB.Where(condition).First(&page).Error
	return page, err
}

func GetPages (c interface {}, p *common.Pagination) ([]Page, error) {
	var pages []Page;

	err := common.DB.Where(c).Order("id desc").Offset(p.GetOffset()).Limit(p.GetLimit()).Find(&pages).Error
	return pages, err
}


func SearchPages(tagname string, p *common.Pagination) ([] Page, error) {
	var pages [] Page
	err := common.DB.Where("tags like ?", "%," + tagname + ",%").Order("id desc").Offset(p.GetOffset()).Limit(p.GetLimit()).Find(&pages).Error
	return pages, err
}

func GetPagesByIDs (ids [] int) ([]Page, error) {
	var pages []Page;
	err := common.DB.Where("id in ?", ids).Find(&pages).Error
	return pages, err
}

type Comment struct {
	ID              int         `gorm:"column:id;primary_key" json:"id"`
	Email           string      `gorm:"column:email" json:"email"`
	Nickname        string      `gorm:"column:nickname" json:"nickname"`
	Content         string      `gorm:"column:content" json:"content"`
	To              * string    `gorm:"column:to" json:"to"`
	CreateTime      time.Time   `gorm:"column:create_time" json:"create_time"`
	IP              string      `gorm:"column:ip" json:"ip"`
	Website         string      `gorm:"column:website" json:"website"`
	PageID          int         `gorm:"column:page_id" json:"page_id"`
	ParentCommentID * int       `gorm:"column:parent_comment_id" json:"parent_comment_id" sql:"type:bigint REFFERENCES comment(id) ON DELETE CASCADE"`
}

// TableName sets the insert table name for this struct type
func (c *Comment) TableName() string {
	return "comment"
}

func GetComment (c interface{}) (Comment, error) {
	var comment Comment
	err := common.DB.Where(c).First(&comment).Error
	return comment, err
}

func GetComments (c interface {}, p *common.Pagination, order string) ([] Comment, error) {
	var comments []Comment

	offset := (p.Page - 1) * p.Size
	err := common.DB.Where(c).Order(order).Offset(offset).Limit(p.Size).Find(&comments).Error

	// aggregate comment
	ids := make([] int, len(comments))
	for i := range(comments) {
		ids[i] = comments[i].PageID
	}
	return comments, err
}

func CreateComment (v *CommentValidator, target *string, ip string, pageID int) * Comment{
	comment := &Comment{
		Email:          v.Email,
		Nickname:       v.Nickname,
		Content:        v.Content,
		To:             target,
		CreateTime:     time.Now(),
		IP:             ip,
		Website:        v.Website,
		PageID:          pageID,
		ParentCommentID: v.CommentID,
	}

	common.DB.Create(comment)
	return comment
}


type Link struct {
	ID          int       `gorm:"column:id;primary_key" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	Href        string    `gorm:"column:href" json:"href"`
	Description string    `gorm:"column:description" json:"description"`
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time"`
	Display     bool      `gorm:"column:display" json:"display"`
}

// TableName sets the insert table name for this struct type
func (l *Link) TableName() string {
	return "link"
}

func GetLink(c interface{}) (Link, error) {
	var link Link
	err :=common.DB.Where(c).First(&link).Error
	return link, err
}

func GetLinks(c interface{}, p *common.Pagination) ([]Link, error) {
	var links [] Link
	err := common.DB.Where(c).Order("id desc").Offset(p.GetOffset()).Limit(p.GetLimit()).Find(&links).Error
	return links, err
}


func GetDistinctTags() []string {
	type Result struct {
		Name string
		count int
	}
	var results []Result
	common.DB.Raw("select name from tag group by name order by count(name) desc").Scan(&results)

	tags := make([]string, len(results))
	for i := range(results) {
		tags[i] = results[i].Name
	}
	return tags
}


type Tag struct {
	ID     int    `gorm:"column:id;primary_key" json:"id"`
	Name   string `gorm:"column:name" json:"name"`
	PageID int    `gorm:"column:page_id" json:"page_id"`
}

// TableName sets the insert table name for this struct type
func (t *Tag) TableName() string {
	return "tag"
}
