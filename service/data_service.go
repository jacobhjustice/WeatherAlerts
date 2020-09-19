package service

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"

	config "github.com/jacobhjustice/WeatherAlerts/model/configuration"
	model "github.com/jacobhjustice/WeatherAlerts/model/data"
)

type IDataService interface {
	GetUsers()
}

type DataService struct {
	IDataService
	Configuration *config.DataConfiguration
}

func (d *DataService) GetUsers() []*model.User {
	// TODO: error handle
	db, _ := d.getDatabaseInstance()
	row, err := db.Query("SELECT UserId, DisplayName, Email, Zipcode from User")
	if err != nil {
		log.Fatal(err)
		return nil
	}

	returnArr := []*model.User{}
	for row.Next() {
		var id int
		var name string
		var email string
		var zip string
		row.Scan(&id, &name, &email, &zip)
		fmt.Println(zip)
		returnArr = append(returnArr, &model.User{
			UserId:      id,
			DisplayName: name,
			Email:       email,
			Zipcode:     zip,
		})
	}
	return returnArr
}

func (d *DataService) getDatabaseInstance() (*sql.DB, error) {
	database, err := sql.Open("sqlite3", d.Configuration.Database)
	if err != nil {
		return nil, err
	}
	return database, nil
}
