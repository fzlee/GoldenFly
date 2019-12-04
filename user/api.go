package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UsersList(c *gin.Context) {
	users, _:= GetUsers(&User{})

	results := make([] *UserResponse, len(users))

	for i := range results{
		result := (&UserSerializer{c, users[i]}).Response()
		results[i] = &result
	}
	c.JSON(http.StatusOK, gin.H{"data": results, "success": true})
}

func UsersMe(c *gin.Context) {
	user, _ := GetUser(&User{ID: 1})
	serializer := UserSerializer{c, user}
	c.JSON(http.StatusOK, gin.H{"data": serializer.Response(), "success": true})
}

func UsersLogin(c *gin.Context) {

}
