package page

import (
	"github.com/gin-gonic/gin"
	"github.com/guregu/null"
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
