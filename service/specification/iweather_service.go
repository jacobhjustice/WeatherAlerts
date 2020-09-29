package specification

type IWeatherService interface {
	GetWeatherForecast(zip string) error
}
