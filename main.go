package main

import (
	"flag"
	"fmt"
	"github.com/gin-contrib/location"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golden_fly/common"
	"golden_fly/config"
	"golden_fly/other"
	"golden_fly/page"
	"golden_fly/user"
)


func initConfig() {
	if err := config.Load("config/config.yaml"); err != nil {
		fmt.Println("Failed to load configuration")
		panic("init config failed")
	}
}

func initDatabase() *gorm.DB {
	db, err := common.InitDB()
	if err != nil {
		fmt.Println("err open databases")
		panic(err)
	}
	return db
}

func initRouters (engine *gin.Engine){
	router := engine.Group("/api")
	user.RegisterRouter(router)
	page.RegisterRouter(router)
	other.RegisterRouter(engine)
}

func initTemplates (engine *gin.Engine) {
	engine.LoadHTMLGlob("templates/*")
}

func initSession (engine *gin.Engine) {
	store := cookie.NewStore([]byte(config.SessionName))
	engine.Use(sessions.Sessions(config.SessionUserKey, store))
}


func main() {
	initConfig()
	db := initDatabase()
	defer db.Close()
	engine := gin.Default()
	engine.Use(location.Default())
	initSession(engine)
	initRouters(engine)
	// initTemplates(engine)

	// command line interface
	var command string
	flag.StringVar(&command, "command", "runserver", "runserver/migrate/createuser")
	flag.Parse()

	if command == "runserver" {
		engine.Run(config.Get().Addr)
	} else if command == "migrate" {
		fmt.Print("makemigrations")
		user.MakeMigrations()
		page.MakeMigrations()
	} else if command == "createuser" {
		user.CreateUserViaCommandLine()
	} else if command == "changepassword" {
		user.ChangePasswordViaCommandLine()
	} else {
	}
}
