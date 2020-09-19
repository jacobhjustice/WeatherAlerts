package service

import (
    "log"
    "net/http"
	"fmt"
	"io/ioutil"
    "github.com/spf13/viper"
)

// Weather service interacts with Open Weather Map API https://openweathermap.org/api
func GetWeatherForecast() {
	fmt.Println("Retrieving Weather...")

    api := viper.GetString("weather.api_key")
    reqStr := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?zip=%s,%s&appid=%s", "35124", "us", api)
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