package user

import (
	"github.com/gin-gonic/gin"
)

type LoginValidator struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (self *LoginValidator) Bind(c *gin.Context) error {
	c.ShouldBindJSON(self)
	return nil
}
