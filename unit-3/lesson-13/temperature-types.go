//Writing function with custom types
package main

import "fmt"

type celsius float64
type kelvin float64

//Convert °K to °C
func kelvinToCelsius (k kelvin) celsius { //Declares a function that accepts one parameter and returns one result
    return celsius(k - 273.15) //Type conversion
}

func main() {
    var k kelvin = 294.0
    c := kelvinToCelsius(k)

    fmt.Printf("%6.2f°K is %6.2f°C\n", k, c)
}
