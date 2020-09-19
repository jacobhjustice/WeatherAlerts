package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	config "github.com/jacobhjustice/WeatherAlerts/model/configuration"
)

type IWeatherService interface {
	GetWeatherForecast()
}

type WeatherService struct {
	IWeatherService
	Configuration *config.WeatherConfiguration
}

// Weather service interacts with Open Weather Map API https://openweathermap.org/api
func (w *WeatherService) GetWeatherForecast(zip string) {
	fmt.Println("Retrieving Weather...")

	reqStr := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?zip=%s,%s&appid=%s", zip, "us", w.Configuration.APIKey)
	resp, err := http.Get(reqStr)
	fmt.Println(reqStr)
	fmt.Println(resp)

	if err != nil {
		fmt.Println("error")
		log.Fatal(err)
	} else {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(body)

	}
}
