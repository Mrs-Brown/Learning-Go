// Three-index slicing example
package main

import "fmt"

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}

	terrestrial := planets[0:4:4] //Length 4 capacity 4
	worlds := append(terrestrial, "Ceres")

	fmt.Println("Planets array", planets)
	fmt.Println("Terrestrial array, three index slicing", terrestrial)
	fmt.Println("Worlds array, three index slicing with append", worlds)

	//Without three index slicing the original array planets will be altered
	terrestrial = planets[0:4] //Length 4, capacity 8
	worlds = append(terrestrial, "Ceres")

	fmt.Println("Planets array", planets)
}
