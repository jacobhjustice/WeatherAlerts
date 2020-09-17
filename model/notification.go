package model

import (
	"github.com/jacobhjustice/WeatherAlerts/model/enum"
)

type Notification struct {
	NotificationId int
	User           User
	Frequency      enum.Frequency
	Alert          enum.Alert
}
