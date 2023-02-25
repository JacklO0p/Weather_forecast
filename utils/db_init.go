package utils

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/JacklO0p/weather_forecast/globals"
	"github.com/JacklO0p/weather_forecast/models"
	"github.com/glebarez/sqlite"
)

func Connect() error {
	var err error
	globals.Db, err = gorm.Open(sqlite.Open("telegram_user.db"), &gorm.Config{})
	if err != nil {
		fmt.Print("Error while connecting to database", err)
		return err
	}

	return nil
}

func MigrateDB() error {
	err := globals.Db.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Print("Error while migrating database", err)
		return err
	}

	return nil
}
