package service

import (
	model "github.com/jacobhjustice/WeatherAlerts/model/configuration"
	"github.com/spf13/viper"
)

type IConfigurationService interface {
	InitializeConfiguration()
	GetCurrentConfiguration()
}

type ConfigurationService struct {
	IConfigurationService
	path      string
	fileName  string
	extension string
}

func (c ConfigurationService) loadConfigurationFile() error {
	viper.AddConfigPath(c.path)
	viper.SetConfigName(c.fileName)
	viper.SetConfigType(c.extension)
	err := viper.ReadInConfig()
	return err
}

func (c ConfigurationService) InitializeConfiguration() (*model.Configuration, error) {
	err := c.loadConfigurationFile()
	if err != nil {
		return nil, err
	}

	return GetCurrentConfiguration(), nil
}

func GetCurrentConfiguration() *model.Configuration {
	return &model.Configuration{
		Email:   loadEmailConfiguration(),
		Data:    loadDataConfiguration(),
		Weather: loadWeatherConfiguration(),
	}
}

func loadEmailConfiguration() *model.EmailConfiguration {
	email := viper.GetString("email.email")
	host := viper.GetString("email.host")
	password := viper.GetString("email.password")
	port := viper.GetString("email.port")

	return &model.EmailConfiguration{
		Email:    email,
		Host:     host,
		Password: password,
		Port:     port,
	}
}

func loadDataConfiguration() *model.DataConfiguration {
	database := viper.GetString("data.database")
	return &model.DataConfiguration{
		Database: database,
	}
}

func loadWeatherConfiguration() *model.WeatherConfiguration {
	api := viper.GetString("weather.api_key")
	return &model.WeatherConfiguration{
		APIKey: api,
	}
}
