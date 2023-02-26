package models

import (
	"fmt"

	"github.com/JacklO0p/weather_forecast/globals"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ChatID    int64  `gorm:"unique" json:"chatid"`
	SendMessage bool   `json:"sendmessage"`
	Location  string `json:"location"`
	Timer int    `json:"timer"`
	ApparentTemperature string `json:"apparenttemperature"`
	RainingHours        int    `json:"raininghours"`
	TotalRain           int    `json:"totalrain"`
}

func (u *User) UpdateLocation(location int) error {
	return globals.Db.Model(u).Update("location", location).Error
}

func (u *User) UpdateTimeFrame(timeFrame string) error {
	return globals.Db.Model(u).Update("timeFrame", timeFrame).Error
}

func GetUserByID(ID int64) (*User, error) {
	var user User

	err := globals.Db.Where("chat_id=?", ID).First(&user).Error
	if err != nil {
		fmt.Print("\nIndex not found")
		return nil, err
	}

	return &user, nil
}

func GetAllUsers() ([]*User, error) {
	var user []*User

	err := globals.Db.Find(&user).Error
	if err != nil {
		fmt.Print("Error while retrieving users", err)
		return nil, err
	}

	return user, nil
}

// Update user if present, if not create it
func UpdateUser(us *User) error {
	var user User

	err := globals.Db.Where("chat_id=?", us.ChatID).FirstOrCreate(&user).Error
	if err != nil {
		fmt.Print("\nIndex not found")
		return err
	}

	err = globals.Db.Model(&user).Updates(us).Error
	if err != nil {
		fmt.Print("\nError while updating user")
		return err
	}

	return nil
}
