// Package weather provides current forecast given a location.
package weather

var (
	// CurrentCondition stores the current condition.
    CurrentCondition string
	// CurrentLocation stores the current location.
    CurrentLocation  string
)

// Forecast returns the expected forecast for the given location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
