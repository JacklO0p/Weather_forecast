package weather

import (
	"time"
)

var Today = time.Now()
var Tomorrow = Today.AddDate(0, 0, 1)

var CurrentDateString = Today.Format("2006-01-02")
var TomorrowDateString = Tomorrow.Format("2006-01-02")

var TimeFrame int = 1 // <= timeframe in minutes