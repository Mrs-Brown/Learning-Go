// A slice of structures []struct is a collection of zero or more values (a slice) where each value is based on a structure instead of a primitive type like float64
package main

import "fmt"

func main() {
	type location struct {
		name string
		lat  float64
		long float64
	}

	locations := []location{ //Declaring a slice of structures
		{name: "Bradbury Landing", lat: -4.5895, long: 137.4417},
		{name: "Columbia Memorial Station", lat: -14.5684, long: 175.472636},
		{name: "Challenger Memorial Station", lat: -1.9462, long: 354.4734}, //Note the trailing comma
	}

	fmt.Println(locations)
}
