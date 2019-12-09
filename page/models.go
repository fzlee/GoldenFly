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

	offset := (p.Page - 1) * p.Size
	err := common.DB.Where(c).Offset(offset).Limit(p.Size).Find(&pages).Error
	return pages, err
}
