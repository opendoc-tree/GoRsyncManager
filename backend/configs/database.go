package configs

import (
	"GoRsyncManager/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("gorsyncmanager.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	err = DB.AutoMigrate(&models.Host{})
	if err != nil {
		panic("failed to migrate database model" + err.Error())
	}
}
