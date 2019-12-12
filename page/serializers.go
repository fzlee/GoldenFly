package page

import (
	"github.com/gin-gonic/gin"
	"time"
)
type PageSerializer struct {
	c *gin.Context
	* Page
}

type PageResponse struct {
    ID				int		 	`json:"id"`
    URL				string			`json:"url"`
    Title			string			`json:"title"`
    ContentDigest 	string			`json:"content_digest,omitempty"`
    Content			string			`json:"content,omitempty"`
    Keywords 		string			`json:"keywords"`
    Metacontent 	string   		`json:metacontent,omitempty`
    CreateTime		* time.Time			`json:"create_time"`
    UpdateTime		* time.Time			`json:"update_time"`
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
		ID:			   self.ID,
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
	ID              int         `json:"id"`
	Email           string      `json:"email,omitempty"`
	Nickname        string      `json:"nickname"`
	Content         string      `json:"content"`
	To              * string       `json:"to"`
	CreateTime      time.Time   `json:"create_time"`
	IP              string      `json:"ip,omitempty"`
	Website         string      `json:"website,omitempty"`
	PageID          int         `json:"page_id"`
	ParentCommentID * int    `json:"parent_comment_id"`

	Page struct {
		Title 		*string		`json:"title"`
		URL         *string		`json:"url"`
	}  `json:"page,omitempty"`

	ParentComment struct {
		Nickname * string
		ID       * int
	}`json:"parent_comment,omitempty"`
}


func (self * CommentSerializer) CommentResponse (IsAdmin bool) *CommentResponse {
	r := &CommentResponse{
		ID:					self.ID,
		Nickname:			self.Nickname,
		Content:			self.Content,
		To:					self.To,
		CreateTime:			self.CreateTime,
	}

	if self.ParentCommentID != nil {
		pComment, err := GetComment(&Comment{ID: *self.ParentCommentID})
		if err != nil {
			r.ParentComment.Nickname = &pComment.Nickname
			r.ParentComment.ID = &pComment.ID
		}
	}

	if IsAdmin {
		r.Website = self.Website
		r.IP = self.IP
		r.Email = self.Email
		r.ID = self.ID
	}

	page, err := GetPage(&Page{ID: self.PageID})
	if err == nil {
		r.Page.Title = &page.Title
		r.Page.URL = &page.URL
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
		r.Page.Title = &page.Title
		r.Page.URL = &page.URL
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


func (self *LinkSerializer) FullResponse () *LinkResponse {
	return &LinkResponse{
		ID: self.ID,
		Name: self.Name,
		Href: self.Href,
		Description: self.Description,
		CreateTime: self.CreateTime,
		Display: self.Display,
	}
}

