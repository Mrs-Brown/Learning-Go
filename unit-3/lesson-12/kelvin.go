//Writing a function
package main

import "fmt"

//Convert °K to °C
func kelvinToCelsius (k float64) float64 { //Declares a function that accepts one parameter and returns one result
    k -= 273.15
    return k
}

func main() {
    kelvin := 294.0
    celsius := kelvinToCelsius(kelvin)
    fmt.Printf("%v°K is %6.2f°C\n", kelvin, celsius)
}
