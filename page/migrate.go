package page

import "golden_fly/common"

func MakeMigrations() {
	common.DB.AutoMigrate(&Page{})
	common.DB.AutoMigrate(&Link{})
	common.DB.AutoMigrate(&Tag{})
	common.DB.AutoMigrate(&Media{})

	if !common.DB.HasTable(&Comment{}) {
		common.DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Comment{})
		common.DB.Model(&Comment{}).AddForeignKey("parent_comment_id", "comment(id)", "CASCADE", "RESTRICT")
	}
}
