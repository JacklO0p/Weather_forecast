package globals

import (
	"time"

	"github.com/go-telegram/bot"
	"gorm.io/gorm"
)

//global variables
var IsProgramStarted bool = false
var Db *gorm.DB
var Bot *bot.Bot

//time variables
var Today = time.Now()
var Tomorrow = Today.AddDate(0, 0, 1)

var CurrentDateString = Today.Format("2006-01-02")
var TomorrowDateString = Tomorrow.Format("2006-01-02")

var TimeFrame int = 1 // <= timer in minutes
