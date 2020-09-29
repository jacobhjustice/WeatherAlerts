package service

import (
	model "github.com/jacobhjustice/WeatherAlerts/model/configuration"
	"github.com/jacobhjustice/WeatherAlerts/model/enum"
	"github.com/spf13/viper"
)

type IConfigurationService interface {
	InitializeConfiguration()
	GetCurrentConfiguration()
}

type ConfigurationService struct {
	IConfigurationService
	Path       string
	FileName   string
	Extension  string
	LogService ILogService
}

// Public

func (c *ConfigurationService) GetCurrentConfiguration() *model.Configuration {
	return &model.Configuration{
		Email:   c.loadEmailConfiguration(),
		Data:    c.loadDataConfiguration(),
		Weather: c.loadWeatherConfiguration(),
	}
}

func (c *ConfigurationService) InitializeConfiguration() (*model.Configuration, error) {
	err := c.loadConfigurationFile()
	if err != nil {
		return nil, err
	}

	return c.GetCurrentConfiguration(), nil
}

// Private

func (c *ConfigurationService) loadConfigurationFile() error {
	c.LogService.Log("Preparing to load configuration file.", enum.INFO)
	viper.AddConfigPath(c.Path)
	viper.SetConfigName(c.FileName)
	viper.SetConfigType(c.Extension)
	err := viper.ReadInConfig()
	return err
}

func (c *ConfigurationService) loadDataConfiguration() *model.DataConfiguration {
	database := viper.GetString("data.database")
	return &model.DataConfiguration{
		Database: database,
	}
}

func (c *ConfigurationService) loadEmailConfiguration() *model.EmailConfiguration {
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

func (c *ConfigurationService) loadWeatherConfiguration() *model.WeatherConfiguration {
	api := viper.GetString("weather.api_key")
	return &model.WeatherConfiguration{
		APIKey: api,
	}
}
