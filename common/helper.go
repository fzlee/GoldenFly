package common

import (
	"github.com/gin-gonic/gin"
	"math/rand"
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

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// A helper function to generate random string
func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}


func MinINT (a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
