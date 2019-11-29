package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UsersList(c *gin.Context) {

}

func UsersRetrieve(c *gin.Context) {

}

func UsersMe(c *gin.Context) {
	user, _ := FindOneUser(&User{ID: 1})
	serializer := UserSerializer{c, user}
	c.JSON(http.StatusOK, gin.H{"data": serializer.Response(), "success": true})
}