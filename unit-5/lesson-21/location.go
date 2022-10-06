// If you need multiple structures with the same fields, you can define a type
package main

import "fmt"

func main() {
	type location struct {
		lat  float64
		long float64
	}

	var spirit location //Spirit rover at Columbia Memorial Station
	spirit.lat = -14.5684
	spirit.long = 175.472636

	var opportunity location //Opportunity rover at Challenger Memorial Station
	opportunity.lat = -1.9462
	opportunity.long = 354.4734

	fmt.Println(spirit, opportunity)
}
