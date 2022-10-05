// Add to array with append function
package main

import "fmt"

func main() {
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"} //Declare slice and Go makes array in background

	dwarfs = append(dwarfs, "Orcus") //Append Orcus to end of array and slice
	fmt.Println(dwarfs)

	//Append is variadic, so you can pass multiple elements
	dwarfs = append(dwarfs, "Salacia", "Quaoar", "Sedna")
	fmt.Println(dwarfs)
}
