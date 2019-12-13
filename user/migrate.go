package user

import "golden_fly/common"

func MakeMigrations () {
	common.DB.AutoMigrate(&User{})
	common.DB.AutoMigrate(&AuthToken{})
}
