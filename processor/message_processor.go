package processor

import (
	"github.com/jacobhjustice/WeatherAlerts/service"
)

type MessageProcessor struct {
	EmailService   service.EmailService
	LogService     service.LogService
	WeatherService service.WeatherService
}
