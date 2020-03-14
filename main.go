package main
import (
    "github.com/jacobhjustice/WeatherAlerts/messager"
    "github.com/jacobhjustice/WeatherAlerts/routing"
    "github.com/spf13/viper"
);

func main() {
    initializeConfigFile()
    messager.Send()
    routing.Route()
}

func initializeConfigFile() {
    viper.AddConfigPath(".")     
    viper.SetConfigName("config")
    viper.SetConfigType("yaml") 
    err := viper.ReadInConfig()
    if err != nil {
        panic("Error: Ensure config.yaml exists in project root.")
    }
}