// Package weather is for weather.
package weather

var (
	//CurrentCondition definition.
	CurrentCondition string
	//CurrentLocation definition.
	CurrentLocation string
)

// Forecast Function that forecasts.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
