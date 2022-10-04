//Example copying an array to a var
package main

import "fmt"

func main() {
    planets := [...]string{
        "Mercury",
        "Venus",
        "Earth",
        "Mars",
        "Jupiter",
        "Saturn",
        "Uranus",
        "Neptune",
    }

    planetsMarkII := planets  //Copies planets array

    planets[2] = "whoops"  //Edit indice of original array

    fmt.Println(planets)  //Print original array
    fmt.Println(planetsMarkII)  //Print copied array
}
