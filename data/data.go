package data

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
)

func getDatabaseName() string {
	return viper.GetString("data.database") + ".db"
}

func read() {
	database, _ := sql.Open("sqlite3", getDatabaseName())
	return database
}

func initializeDatabase() sql.DB {
	database, _ := sql.Open("sqlite3", getDatabaseName())
	return database
}
