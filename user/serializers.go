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
	UID       string    `json:"uid,omitempty"`
	Username  string    `json:"username,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Role      string    `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (self *UserSerializer) Response () *UserResponse {
	return &UserResponse{
		UID: self.UID,
		Username: self.Username,
		Nickname: self.Nickname,
		Role: self.Role,
		CreatedAt: self.CreatedAt,
		UpdatedAt: self.UpdatedAt,
	}
}

func (self *UserSerializer) LoginResponse() *UserResponse {
	return &UserResponse{
		UID:       self.UID,
		Nickname:  self.Nickname,
		Role:      self.Role,
	}
}

