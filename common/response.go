package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


type GeneralResponse struct {
	success bool
	data    *gin.H
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

func ResponseWithCode(c * gin.Context, code int) {
	c.JSON(http.StatusOK, gin.H{
		"success": false,
		"error": gin.H{
			"code": code,
			"message": ErrorCodes[code],
		},
	})
}

func ResponseWithPanic(c *gin.Context, err error) {
	c.JSON(http.StatusBadGateway, gin.H{
		"success": false,
	})
}

func ResponseWithValidation(c * gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"error": err.Error(),
	})
}

func ResponseWithData (c *gin.Context, data * gin.H) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": data,
	})
}
