package specification

import (
	"github.com/jacobhjustice/WeatherAlerts/model/configuration"
)

type IServiceInitializer interface {
	InitializeConfigurationService() IConfigurationService
	InitializeDataService(config *configuration.DataConfiguration) IDataService
	InitializeEmailService(config *configuration.EmailConfiguration) IEmailService
	InitializeLogService() ILogService
	InitializeWeatherService(config *configuration.WeatherConfiguration) IWeatherService
}
