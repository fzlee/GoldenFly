package user

import (
	"github.com/gin-gonic/gin"
	"golden_fly/common"
	"net/http"
)

func UsersListView(c *gin.Context) {
	users, _ := GetUsers(&User{})

	results := make([]*UserResponse, len(users))

	for i := range results {
		results[i] = (&UserSerializer{c, users[i]}).Response()
	}
	c.JSON(http.StatusOK, gin.H{"data": results, "success": true})
}

func UsersMeView(c *gin.Context) {
	user, _ := GetUser(&User{ID: 1})
	serializer := UserSerializer{c, user}
	c.JSON(http.StatusOK, gin.H{"data": serializer.Response(), "success": true})
}

func LoginView(c *gin.Context) {

	loginValidator := LoginValidator{}
	if err := loginValidator.Bind(c); err != nil {
		common.ResponseWithCode(c, common.CodeParameterMissing)
		return
	}
	user, err := GetUser(&User{Username: loginValidator.Username})

	if err != nil {
		common.ResponseWithCode(c, common.CodeLoginFailed)
		return
	}

	if !user.CheckPassword(loginValidator.Password) {
		common.ResponseWithCode(c, common.CodeLoginFailed)
		return
	}
	serializer := UserSerializer{c, user}

	token := user.GetOrExtendToken()
	WriteCredentialToCookie(c, &user, &token)
	c.JSON(http.StatusOK, gin.H{"data": serializer.LoginResponse(), "success": true})
}
