// Repurposing maps with sets
package main

import (
	"fmt"
	"sort"
)

func main() {
	var temperatures = []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	set := make(map[float64]bool) //Makes a map with Boolean values
	for _, t := range temperatures {
		set[t] = true //For every element, assign true
	}

	if set[-28.0] {
		fmt.Println("set member")
	}
	fmt.Println(set) //Prints map[-31:true -29:true -23:true -33:true -28:true 32:true]

	//Before they can be sorted, the temperatures must be converted back to a slice, because maps have an arbitrary order

	unique := make([]float64, 0, len(set))
	for t := range set {
		unique = append(unique, t)
	}
	sort.Float64s(unique)

	fmt.Println(unique) //Prints [-33 -31 -29 -28 -23 32]
}
