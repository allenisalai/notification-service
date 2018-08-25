package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/allenisalai/notification-service/internal/model"
)

var db *gorm.DB

func init() {
	d, err := gorm.Open("postgres", "host=localhost port=5432 user=alai dbname=notifications sslmode=disable")
	if err != nil {
		panic("failed to connect database:" + err.Error())
	}
	db = d
	initializeTables()
}

func GetDb() *gorm.DB {
	return db
}

func initializeTables() {
	db.AutoMigrate(&model.NotificationType{})
}