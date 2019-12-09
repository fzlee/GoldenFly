package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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
	router := engine.Group("/api")
	user.RegisterRouter(router)
	page.RegisterRouter(router)
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
	initSession(engine)
	initRouters(engine)
	engine.Run(config.Get().Addr)
}
