// Package weather provides current weather information for a specified city.
package weather

// CurrentCondition is the current weather condition.
var CurrentCondition string
// CurrentLocation is the current location for which the weather is being described.
var CurrentLocation string

// Forecast returns the current weather conditions for a specified city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
