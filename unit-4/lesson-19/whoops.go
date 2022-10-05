// Maps aren't copied but instead point to the same underlying data
package main

import "fmt"

func main() {
	planets := map[string]string{
		"Earth": "Sector ZZ9",
		"Mars":  "Sector ZZ9",
	}

	planetsMarkII := planets
	planets["Earth"] = "whoops"

	fmt.Println(planets)       //Prints Earth:whoops Mars:Sector ZZ9
	fmt.Println(planetsMarkII) //Prints Earth:whoops Mars:Sector ZZ9

	delete(planets, "Earth")   //Removes Earth from map
	fmt.Println(planetsMarkII) //Prints Mars:Sector ZZ9
}
