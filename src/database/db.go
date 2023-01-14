package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"project/src/models"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=db user=root password=root dbname=ambassador port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect with database!")
	}

}
func AutoMigrate() {
	DB.AutoMigrate(models.User{})
}
