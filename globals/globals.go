package globals

import (
	"time"

	"github.com/go-telegram/bot"
	"gorm.io/gorm"
)

var IsProgramStarted bool = false
var Db *gorm.DB
var Bot *bot.Bot

var CurrentDateString = time.Now().Format("2006-01-02")
var TomorrowDateString = time.Now().AddDate(0, 0, 1).Format("2006-01-02")

var Timer int = 1 // <= timer in minutes

var TimerDuration = time.Duration(Timer) * time.Minute