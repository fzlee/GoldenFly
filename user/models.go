package user

import (
	"database/sql"
	"github.com/alexandrevicenzi/unchained"
	"github.com/guregu/null"
	"time"

	"golden_fly/common"
	"golden_fly/config"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

type User struct {
	ID        uint       `gorm:"column:id;primary_key" json:"id"`
	Password  string    `gorm:"column:password" json:"password"`
	LastLogin time.Time `gorm:"column:last_login" json:"last_login"`
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

func GetUser (condition interface{}) (User, error) {
	var user User
	err := common.DB.Where(condition).First(&user).Error
	return user, err
}

func GetUsers(condition interface{}) ([]User, error) {
	var users []User
	error := common.DB.Where(condition).Find(&users).Error
	return users, error
}

func (self *User) CheckPassword(password string) bool {
	valid, _ := unchained.CheckPassword(password, self.Password)

	if valid {
		return true
	}
	return false
}

func (self *User) HashPassword(password string) string {
	hash, err := unchained.MakePassword(password, unchained.GetRandomString(12), "default")

	if err == nil {
		return hash
	} else {
		return ""
	}
}


func (self *User) SetPassword (password string) error {
	self.Password = self.HashPassword(password)
	return common.DB.Save(self).Error
}

func (self *User) GetOrExtendToken () AuthToken {
	token, err := GetToken(&AuthToken{UID: self.UID})
	if err != nil  {
		return *createTokenFor(self.UID)
	} else if token.IsGoingToExpired() {
		token.extendToken()
		return token
	} else {
		return token
	}
}

type AuthToken struct {
	Key       string    `gorm:"column:key;primary_key" json:"key"`
	UID       string    `gorm:"column:uid" json:"uid"`
	Created   time.Time `gorm:"column:created" json:"created"`
	ExpiredAt time.Time `gorm:"column:expired_at" json:"expired_at"`
}

// TableName sets the insert table name for this struct type
func (a *AuthToken) TableName() string {
	return "auth_token"
}

func GetToken(condition interface{}) (AuthToken, error) {
	var token AuthToken
	err := common.DB.Where(condition).First(&token).Error
	return token, err
}

func createTokenFor(uid string) *AuthToken {
	key := common.RandomString(20)
	now := time.Now()
	conf := config.Get()

	token := &AuthToken{
		Key:       key,
		UID:       uid,
		Created:   now,
		ExpiredAt: now.AddDate(0, 0, conf.TokenExpiredDays),
	}
	common.DB.Create(token)
	return token
}

func (self *AuthToken) extendToken(){
	now := time.Now()
	conf := config.Get()
	common.DB.Delete(&self)
	self.Key = common.RandomString(20)
	self.ExpiredAt = now.AddDate(0, 0, conf.TokenExpiredDays)
	common.DB.Save(&self)
}

func (self *AuthToken) HasExpired () bool {
	now := time.Now()
	return self.ExpiredAt.Before(now)
}

func (self *AuthToken) IsGoingToExpired() bool {
	conf := config.Get()
	then := time.Now()
	then.AddDate(0, 0, conf.TokenRefreshDays)
	return self.ExpiredAt.Before(then)
}


func ValidateToken(key string) (User, error){
	token, err := GetToken(&AuthToken{Key: key})
	if err == nil && !token.HasExpired(){
		var user, err = GetUser(&User{UID: token.UID})
		return user, err
	}
	return User{}, err
}
