package main

import (
	"app/databases"
	"app/models"
	"github.com/sirupsen/logrus"
)

func main() {
	db, err := databases.Connect() //*gorm.DBが返却される
	defer db.Close()

	if err != nil {
		logrus.Fatal(err)
	}

	db.Debug().AutoMigrate(&models.User{}) //AutoMigrate(モデル）でマイグレーションする
	db.Debug().AutoMigrate(&models.Favorite{})
}
