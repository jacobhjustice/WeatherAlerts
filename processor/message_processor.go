package processor

import (
	"github.com/jacobhjustice/WeatherAlerts/service/specification"
)

type MessageProcessor struct {
	EmailService   *specification.IEmailService
	LogService     *specification.ILogService
	WeatherService *specification.IWeatherService
}
