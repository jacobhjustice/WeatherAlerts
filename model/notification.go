package model

import (
    "../model/enum"
)

type Notification struct {
	NotificationId
	User User
	Frequency Frequency
	Alert Alter
}