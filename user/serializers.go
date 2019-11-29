package user

import (
	"github.com/gin-gonic/gin"
	"time"
)

type UserSerializer struct {
	C *gin.Context
	User
}

type UserResponse struct {
	ID        uint      `json:"-"`
	UID       string    `gorm:"column:uid" json:"uid"`
	Username  string    `gorm:"column:username" json:"username"`
	Nickname  string    `gorm:"column:nickname" json:"nickname"`
	Role      string    `gorm:"column:role" json:"role"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (self *UserSerializer) Response () UserResponse {
	return UserResponse{
		ID: self.ID,
		UID: self.UID,
		Username: self.Username,
		Nickname: self.Nickname,
		Role: self.Role,
		CreatedAt: self.CreatedAt,
		UpdatedAt: self.UpdatedAt,
	}
}