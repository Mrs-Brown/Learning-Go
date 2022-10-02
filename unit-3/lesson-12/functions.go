//Writing multiple functions
package main

import "fmt"

//Convert °K to °C
func kelvinToCelsius (k float64) float64 { //Declares a function that accepts one parameter and returns one result
    k -= 273.15
    return k
}

//Convert °C to °F
func celsiusToFahrenheit (c float64) float64 {
    c = (c * 9.0 / 5.0) + 32.0
    return c
}

//Convert °K to °F
func kelvinToFahrenheit (k float64) float64 {
    k = 1.8 * (k - 273.0) + 32.0
    return k 
}

func main() {
    kelvin := 233.0
    celsius := kelvinToCelsius(kelvin)
    fahrenheitFromCelsius := celsiusToFahrenheit(celsius)
    fahrenheitFromKelvin := kelvinToFahrenheit(kelvin)

    fmt.Printf("%6.2f°K is %6.2f°C\n", kelvin, celsius)
    fmt.Printf("%6.2f°C is %6.2f°F\n",celsius, fahrenheitFromCelsius)
    fmt.Printf("%6.2f°K is %6.2f°F\n", kelvin, fahrenheitFromKelvin)
}
