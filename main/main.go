package main

import (
	"fmt"

	"github.com/jacobhjustice/WeatherAlerts/model/configuration"
	"github.com/jacobhjustice/WeatherAlerts/service"
)

func main() {
	config := initializeConfigFile()
	dataService := service.DataService{
		Configuration: config.Data,
	}

	result := dataService.GetUsers()
	fmt.Println(result[0])
	// messager.Send()
	// routing.Route()
	// result := data.GetUsers()
	// fmt.Println(result[0])
	// //weather.GetWeatherForecast()
}

func initializeConfigurationService() *service.ConfigurationService {
	return service.ConfigurationService{
		Path:      ".",
		FileName:  "config",
		Extension: "yaml",
		Log:       &service.LogService{},
	}

	
	return config

}

funcService