package user

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/ssh/terminal"
	"golden_fly/common"
	"golden_fly/config"
	"net/http"
	"net/url"
	"os"
	"strings"
	"syscall"
	"time"
)

func WriteCredentialToCookie (c *gin.Context, user *User, token *AuthToken) {
	conf := config.Get()
	http.SetCookie(c.Writer, &http.Cookie{
		Name:       conf.CookieToken,
		Value:      token.Key,
		Expires:    token.ExpiredAt,
		Path:       "/",
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
		Path:       "/",
	})
}


func ChangePasswordViaCommandLine() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Username: ")
	Username, _ := reader.ReadString('\n')

	fmt.Print("Enter Password: ")
	bPassword1, _ := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Print("Repeat Password: ")
	bPassword2 , _ := terminal.ReadPassword(int(syscall.Stdin))

	if string(bPassword1) != string(bPassword2) {
		fmt.Println("Password confirmation failed")
		return
	}

	password := string(bPassword1)

	user, err := GetUser(&User{Username: Username})
	if err != nil {
		fmt.Println("User not found")
		return
	}

	user.SetPassword(password)
	fmt.Println("Done")
}


func CreateUserViaCommandLine () {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Username: ")
	Username, _ := reader.ReadString('\n')
	Username = strings.Trim(Username, "\n")

	fmt.Println("Enter Password: ")
	bPassword1, _ := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println("Repeat Password: ")
	bPassword2 , _ := terminal.ReadPassword(int(syscall.Stdin))

	if string(bPassword1) != string(bPassword2) {
		fmt.Println("Password confirmation failed")
		return
	}

	password := string(bPassword1)
	user, err := GetUser(&User{Username: Username})
	if err == nil {
		fmt.Println("Username is in use")
	}

	now := time.Now()
	user = User{
		LastLogin: now,
		UID:   common.RandomString(12),
		Username:  Username,
		Activated: 1,
		CreatedAt: now,
		UpdatedAt: now,
	}
	user.SetPassword(password)
	common.DB.Save(&user)
	fmt.Println("Done")
}
