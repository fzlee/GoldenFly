package page

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Page struct {
	gorm.Model
	Url string
	title string
	content string
}