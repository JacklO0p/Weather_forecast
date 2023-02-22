package listener

import (
	"strings"

	"github.com/JacklO0p/weather_forecast/models"
)

var CommandsList [10]models.Command

func isCommandPresent(com string) bool {

	for index, _ := range CommandsList {
		if strings.Contains(CommandsList[index].Com, com) {
			return true
		}
	}

	return false
}

func Inizializer() {
	CommandsList[0].Com = "/start"
	CommandsList[1].Com = "/location <city name>"
	CommandsList[2].Com = "/meteo"
	CommandsList[3].Com = "/timeframe"
	CommandsList[4].Com = "/newtimeframe <time frame in minutes>"
	CommandsList[5].Com = "/help"
	CommandsList[6].Com = "/stop"
}

func ListOfCommands() (list string) {

	for index, _ := range CommandsList {
		list += CommandsList[index].Com + "\n"
	}

	return list
}
