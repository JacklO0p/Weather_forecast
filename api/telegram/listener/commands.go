package listener

import (
	"strings"

	"github.com/JacklO0p/weather_forecast/models"
)

var CommandsList [10]models.Command

func isCommandPresent(com string) bool {

	for index := range CommandsList {
		if strings.Contains(CommandsList[index].Com, com) {
			return true
		}
	}

	return false
}

func Inizializer() {
	CommandsList[0].Com = "/start"
	CommandsList[1].Com = "/location <city name>"
	CommandsList[2].Com = "/report"
	CommandsList[3].Com = "/timer"
	CommandsList[4].Com = "/newtimer <time frame in minutes>"
	CommandsList[5].Com = "/help"
}

func ListOfCommands() (list string) {

	for index := range CommandsList {
		list += CommandsList[index].Com + "\n"
	}

	return list
}
