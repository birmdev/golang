// Package weather forecasts the current weather condition of various cities in the country Goblinocus.
package weather

var (
	// CurrentCondition is a description of the current weather condition.
	CurrentCondition string
	// CurrentLocation is the name of the current location/city.
	CurrentLocation string
)

// Forecast provides the current weather condition of the current city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
