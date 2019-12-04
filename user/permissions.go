package user

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golden_fly/common"
	"golden_fly/config"
	"net/http"
)

func AdminRequired(c *gin.Context){
	session := sessions.Default(c)
	user := session.Get(config.SessionUserKey)
	if user != nil {
		// Abort the request with the appropriate error code
		common.AbortWithCode(c, http.StatusForbidden, common.CodePermissionDenied)
	}
	// Continue down the chain to handler etc
	c.Next()
}
