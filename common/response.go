package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorResponseSerializer struct {
	C *gin.Context
	Code	uint
	Message string
}

func AbortWithCode(c *gin.Context, httpCode int, code int) {
	c.AbortWithStatusJSON(httpCode, gin.H{
		"success": false,
		"error": gin.H{
			"code": code,
			"message": ErrorCodes[code],
		},
	})
}

func ResponseWithCode(c *gin.Context, code int) {
	c.JSON(http.StatusOK, gin.H{
		"success": false,
		"error": gin.H{
			"code": code,
			"message": ErrorCodes[code],
		},
	})
}
