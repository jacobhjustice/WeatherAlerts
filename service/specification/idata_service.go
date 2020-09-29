package specification

import model "github.com/jacobhjustice/WeatherAlerts/model/data"

type IDataService interface {
	GetUsers() ([]*model.User, error)
}
