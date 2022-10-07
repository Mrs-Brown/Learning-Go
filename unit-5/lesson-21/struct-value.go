// Structs are copied
package main

import "fmt"

func main() {
	type location struct {
		lat, long float64
	}

	bradbury := location{-4.5895, 137.4417} //Bradbury Landing
	curiosity := bradbury                   //Curiosity is at same location

	curiosity.long += 0.0106 //Curiosity moves east to Yellowknife Bay

	fmt.Println(bradbury, curiosity) //Prints the independent lat and long values
}
