// Weather service interacts with Open Weather Map API https://openweathermap.org/api

package implementation

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	config "github.com/jacobhjustice/WeatherAlerts/model/configuration"
	"github.com/jacobhjustice/WeatherAlerts/service/specification"
)

type WeatherService struct {
	specification.IWeatherService
	Configuration *config.WeatherConfiguration

	LogService *specification.ILogService
}

func (w WeatherService) GetWeatherForecast(zip string) error {
	fmt.Println("Retrieving Weather...")

	reqStr := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?zip=%s,%s&appid=%s", zip, "us", w.Configuration.APIKey)
	resp, err := http.Get(reqStr)
	fmt.Println(reqStr)
	fmt.Println(resp)

	if err != nil {
		fmt.Println("error")
		log.Fatal(err)
		return err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(body)
	return err

}
