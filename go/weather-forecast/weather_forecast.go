// Package weather forecasts the current weather conditions of various cities in Goblinocus.
package weather

// CurrentCondition variable holds a string that describes the current condition of weather.
var CurrentCondition string

// CurrentLocation variable holds a string naming the current location.
var CurrentLocation string

// Forecast takes currentlocation and currentweather and returns them in a formatted string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
