//Modifying a slice
package main

import "fmt"

//Modify the length of the planets slice
//by the use of a pointer
func reclassify(planets *[string]) { 
	*planets = (*planets)[0:8]
}

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
		"Pluto",
	}
	reclassify(&planets)
	fmt.Println(planets) 
}