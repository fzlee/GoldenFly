package main

import (
	"flag"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golden_fly/common"
	"golden_fly/config"
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
	conf := config.Get()
	router := engine.Group("/api")
	user.RegisterRouter(router)
	page.RegisterRouter(router)

	// sitemap
	router.GET("/sitemap.tmpl", page.GenerateSitemapView)
	// static folder
	engine.StaticFS("/media", gin.Dir(conf.MediaFolder, false))
}

func initTemplates (engine *gin.Engine) {
	engine.LoadHTMLGlob("templates/*")
	page.RegisterTemplateViews(engine)
}

func initSession (engine *gin.Engine) {
	store := cookie.NewStore([]byte(config.SessionName))
	engine.Use(sessions.Sessions(config.SessionUserKey, store))
}

func initCLI () {

}

func main() {
	initConfig()
	db := initDatabase()
	defer db.Close()
	engine := gin.Default()
	initSession(engine)
	initRouters(engine)
	initTemplates(engine)

	// command line interface
	var command string
	flag.StringVar(&command, "command", "xxx", "runserver/migrate/createuser")
	flag.Parse()
	fmt.Println(command)
	fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxx")

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
	}
}
