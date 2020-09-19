package service

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"

	model "github.com/jacobhjustice/WeatherAlerts/model/data"
)

func getDatabaseName() string {
	return viper.GetString("data.database")
}

func GetUsers() []*model.User {
	db := initializeDatabase()
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

func initializeDatabase() *sql.DB {
	database, _ := sql.Open("sqlite3", getDatabaseName())
	return database
}
