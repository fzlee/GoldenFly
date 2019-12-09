package common

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func ParsePageAndSize(c *gin.Context) Pagination {
	page := c.DefaultQuery("page", "1")
	size := c.DefaultQuery("size", "12")

	intPage, error := strconv.Atoi(page)
	if error != nil {
		intPage = 1
	}

	intSize, error := strconv.Atoi(size)
	if error != nil {
		intSize = 12
	}
	if intSize > 100 {
		intSize = 100
	}

	return Pagination{intPage, intSize}
}
