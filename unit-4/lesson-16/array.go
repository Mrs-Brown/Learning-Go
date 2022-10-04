//Using arrays
package main 

import "fmt"

func main() {
    var planets [8]string  //Planets array contains 8 string elements

    planets[0] = "Mercury"  //Assign a planet at index 0
    planets[1] = "Venus"
    planets[2] = "Earth"
    planets[3] = "Mars"
    planets[4] = "Jupiter"
    planets[5] = "Saturn"
    planets[6] = "Uranus"
    planets[7] = "Neptune"

    earth := planets[2]  //Retrieves the planet at index 2
    fmt.Println(earth)  //Prints Earth
}
