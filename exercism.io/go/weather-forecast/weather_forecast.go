// Package weather provides weather forecast information for a given city.
package weather

// CurrentCondition returns the current weather condition for a given city.
var CurrentCondition string
// CurrentLocation returns the current location for a given city.
var CurrentLocation string

// Forecast returns the weather forecast for a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
