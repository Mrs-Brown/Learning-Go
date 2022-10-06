// Introducing structures
package main

import "fmt"

func main() {
	var rocketCity struct {
		lat  float64 //latitude
		long float64 //longitude
	}

	rocketCity.lat = 34.7111
	rocketCity.long = 86.6539 //Latitude and longitude for space and rocket center, Huntsville AL

	fmt.Println(rocketCity.lat, rocketCity.long)
	fmt.Println(rocketCity) //Both print statements show same results
}
