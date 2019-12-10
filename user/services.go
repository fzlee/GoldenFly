package user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golden_fly/config"
	"net/http"
	"net/url"
	"strings"
)

func WriteCredentialToCookie (c *gin.Context, user *User, token *AuthToken) {
	conf := config.Get()
	http.SetCookie(c.Writer, &http.Cookie{
		Name:       conf.CookieToken,
		Value:      token.Key,
		Expires:    token.ExpiredAt,
	})

	cookieUser := (&UserSerializer{c, *user}).LoginResponse()
	cookieData, _ := json.Marshal(cookieUser)
	cookieData2 := string(cookieData)
	cookieData3 := strings.ReplaceAll(cookieData2, " ","%20")
	cookieData4 := url.QueryEscape(cookieData3)


	http.SetCookie(c.Writer, &http.Cookie{
		Name:       conf.CookieUser,
		Value:      cookieData4,
		Expires:    token.ExpiredAt,
	})
}
