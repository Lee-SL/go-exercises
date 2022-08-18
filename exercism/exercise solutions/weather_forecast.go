// Package weather d.
package weather

// CurrentCondition d.
var CurrentCondition string

// CurrentLocation d.
var CurrentLocation string

// Forecast d.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
