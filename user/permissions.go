package user

import (
	"github.com/gin-gonic/gin"
	"golden_fly/common"
	"golden_fly/config"
	"net/http"
)

func AdminRequired(c *gin.Context){
	conf := config.Get()
	key, err1 := c.Cookie(conf.CookieToken)
	if err1 != nil {
		common.AbortWithCode(c, http.StatusForbidden, common.CodePermissionDenied)
		return
	}
	var _ , err2 = ValidateToken(key)
	if err2 != nil {
		common.AbortWithCode(c, http.StatusForbidden, common.CodePermissionDenied)
		return
	}
	c.Next()
}
