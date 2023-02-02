// Package weather contains forecasts.
package weather

// CurrentCondition : will check the weather condition in a given city.
var CurrentCondition string

// CurrentLocation : will be the location of the weather being checked.
var CurrentLocation string

// Forecast : will return the value of CurrentLocation + the CurrentCondition of weather within that location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
