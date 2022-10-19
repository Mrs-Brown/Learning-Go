//When a var is declared as a function type, its value is nil by default
//This program sorts a slice of strings with the less function.
//If nil is passed for the less argument, it defaults to an alphabetical sort.
package main

import (
	"fmt"
	"sort"
)

func sortStrings(s []string, less func(i, j int) bool) {
	if less == nil {
		less = func(i, j int) bool { return s[i] < s[j] }
	}
	sort.Slice(s, less)
}

func main() {
	food := []string{"onion", "carrot", "celery"}
	sortStrings(food, nil)
	fmt.Println(food)  //Prints [carrot celery onion]
}