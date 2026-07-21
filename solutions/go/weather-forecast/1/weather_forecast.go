//Package weather provides a tool for getting a weather forecast.
package weather



var (
    //CurrentCondition represents the prevailing weather in that given location.
	CurrentCondition string
    //CurrentLocation represents the place with that given weather condition.
	CurrentLocation  string
)

//Forecast returns a string showing the current weather condition of a given place or city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
