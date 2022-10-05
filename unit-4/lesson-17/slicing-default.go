//Default values when slicing
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

    terrestrial := planets[:4]  //default from begining to indice 4
    gasGiants := planets[4:6]
    iceGiants := planets[6:]  //default from indice 6 to end

    allPlanets := planets[:]  //default beginning to end

    fmt.Println(terrestrial, gasGiants, iceGiants, allPlanets)
}
