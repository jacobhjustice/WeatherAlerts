package implementation

import (
	"github.com/jacobhjustice/WeatherAlerts/model/configuration"
	"github.com/jacobhjustice/WeatherAlerts/service/specification"
)

type ServiceInitializer struct {
	specification.IServiceInitializer
}

func (i *ServiceInitializer) InitializeConfigurationService() specification.IConfigurationService {
	log := i.InitializeLogService()
	return ConfigurationService{
		Path:       ".",
		FileName:   "config",
		Extension:  "yaml",
		LogService: &log,
	}
}

func (i *ServiceInitializer) InitializeDataService(config *configuration.DataConfiguration) specification.IDataService {
	log := i.InitializeLogService()
	return DataService{
		Configuration: config,
		LogService:    &log,
	}
}

func (i *ServiceInitializer) InitializeEmailService(config *configuration.EmailConfiguration) specification.IEmailService {
	log := i.InitializeLogService()
	return EmailService{
		Configuration: config,
		LogService:    &log,
	}
}

func (i *ServiceInitializer) InitializeLogService() specification.ILogService {
	return LogService{}
}

func (i *ServiceInitializer) InitializeWeatherService(config *configuration.WeatherConfiguration) specification.IWeatherService {
	log := i.InitializeLogService()
	return WeatherService{
		Configuration: config,
		LogService:    &log,
	}
}
