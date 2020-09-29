package specification

import model "github.com/jacobhjustice/WeatherAlerts/model/configuration"

type IConfigurationService interface {
	InitializeConfiguration() *model.Configuration
	GetCurrentConfiguration() (*model.Configuration, error)
}
