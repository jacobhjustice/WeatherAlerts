package specification

import model "github.com/jacobhjustice/WeatherAlerts/model/configuration"

type IConfigurationService interface {
	GetCurrentConfiguration() *model.Configuration
	InitializeConfiguration() (*model.Configuration, error)
}
