package specification

import "github.com/jacobhjustice/WeatherAlerts/model/enum"

type ILogService interface {
	Log(message string, severity enum.LogSeverity)
}
