package user

import (
	"database/sql"
	"time"

	"github.com/guregu/null"

	"golden_fly/common"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

type User struct {
	ID        uint       `gorm:"column:id;primary_key" json:"id"`
	Password  string    `gorm:"column:password" json:"password"`
	LastLogin null.Time `gorm:"column:last_login" json:"last_login"`
	UID       string    `gorm:"column:uid" json:"uid"`
	Username  string    `gorm:"column:username" json:"username"`
	Email     string    `gorm:"column:email" json:"email"`
	Nickname  string    `gorm:"column:nickname" json:"nickname"`
	Role      string    `gorm:"column:role" json:"role"`
	Avatar    string    `gorm:"column:avatar" json:"avatar"`
	Activated int       `gorm:"column:activated" json:"activated"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "user"
}

func FindOneUser(condition interface{}) (User, error) {
	var user User
	err := common.DB.Where(condition).First(&user).Error
	return user, err
}