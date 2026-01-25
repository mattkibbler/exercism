// Package weather provides stuff about weather.
package weather

var (
	// CurrentCondition represents the current condition
	CurrentCondition string
	// CurrentLocation represents the current location
	CurrentLocation string
)

// Forecast returns the forecast for a city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
