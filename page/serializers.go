package page

import (
	"github.com/gin-gonic/gin"
	"github.com/guregu/null"
	"golden_fly/common"
	"time"
)
type PageSerializer struct {
	c *gin.Context
	* Page
}

type PageResponse struct {
    ID				uint		 	`json:"_"`
    URL				string			`json:"url"`
    Title			string			`json:"title"`
    ContentDigest 	string			`json:"content_digest,omitempty"`
    Content			string			`json:"content,omitempty"`
    Keywords 		string			`json:"keywords"`
    Metacontent 	string   		`json:metacontent,omitempty`
    CreateTime		time.Time			`json:"create_time"`
    UpdateTime		time.Time			`json:"update_time"`
    NeedKey			bool			`json:"need_key"`
    Password		string			`json:"password,omitempty"`
    Tags			string			`json:"tags"`
    Editor			string			`json:"editor"`
    AllowVisit		bool			`json:"allow_visit"`
    AllowComment	bool			`json:"allow_comment"`
    IsOriginal		bool			`json:"is_original"`
    NumLookup		int				`json:"num_lookup"`
    HTML 			string			`json:"html,omitempty"`
    Preview			string			`json:"preview,omitempty"`
}

func (self * PageSerializer) FullResponse() *PageResponse {
	return &PageResponse{
		URL:           self.URL,
		Title:         self.Title,
		ContentDigest: self.ContentDigest,
		Content:       self.Content,
		Keywords:      self.Keywords,
		Metacontent:   self.MetaContent,
		CreateTime:    self.CreateTime,
		UpdateTime:    self.UpdateTime,
		NeedKey:       self.NeedKey,
		Tags:          self.Tags,
		Editor:        self.Editor,
		AllowVisit:    self.AllowVisit,
		AllowComment:  self.AllowComment,
		IsOriginal:    self.IsOriginal,
		NumLookup:     self.NumLookup,
		HTML:          self.HTML,
	}
}

func (self * PageSerializer) PreviewResponse() *PageResponse {
	r :=  &PageResponse{
		URL:           self.URL,
		Title:         self.Title,
		CreateTime:    self.CreateTime,
		NeedKey:       self.NeedKey,
		Tags:          self.Tags,
	}

	if !self.NeedKey {
		r.Preview = self.ContentDigest
	}
	return r
}

func (self * PageSerializer) MetaResponse() *PageResponse {
	r :=  &PageResponse{
		URL:           self.URL,
		Title:         self.Title,
		Tags:          self.Tags,
		CreateTime:    self.CreateTime,
		NeedKey:       self.NeedKey,
		AllowComment:  self.AllowComment,
		IsOriginal:	   self.IsOriginal,
		Content:       self.Content,
	}

	r.Content = self.Content
	return r
}


type CommentSerializer struct {
	c *gin.Context
	* Comment
}

type CommentResponse struct {
	ID              int         `json:"_"`
	Email           string      `json:"email"`
	Nickname        string      `json:"nickname"`
	Content         string      `json:"content"`
	To              null.String `json:"to"`
	CreateTime      time.Time   `json:"create_time"`
	IP              string      `json:"ip"`
	Website         string      `json:"website"`
	PageID          int         `json:"page_id"`
	ParentCommentID null.Int    `json:"parent_comment_id"`

	Page struct {
		Title 		string		`json:"title"`
		URL         string		`json:"url"`
	}  `json:"page,omitempty"`
}


func (self * CommentSerializer) CommentResponse (IsAdmin bool) *CommentResponse {
	r := &CommentResponse{
		Nickname:			self.Nickname,
		Content:			self.Content,
		To:					self.To,
		CreateTime:			self.CreateTime,
	}

	if IsAdmin {
		r.Website = self.Website
		r.IP = self.IP
		r.Email = self.Email
	}
	return r
}


func (self * CommentSerializer) SidebarCommentResponse () *CommentResponse {

	r := &CommentResponse{
		Nickname:        self.Nickname,
		Content:         self.Content,
		To:              self.To,
		CreateTime:		 self.CreateTime,
	}

	page, err := GetPage(&Page{ID: self.PageID})
	if err == nil {
		r.Page.Title = page.Title
		r.Page.URL = page.URL
	}
	return r
}

type LinkSerializer struct {
	c *gin.Context
	* Link
}

type LinkResponse struct {
	ID          int       `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Href        string    `json:"href,omitempty"`
	Description string    `json:"description,omitempty"`
	CreateTime  time.Time `json:"create_time,oitempty"`
	Display     bool      `json:"display,omitempty"`
}

func (self *LinkSerializer) getSidebarResponse () *LinkResponse {
	return &LinkResponse{
		Name: self.Name,
		Href: self.Href,
		Description: self.Description,
	}
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
