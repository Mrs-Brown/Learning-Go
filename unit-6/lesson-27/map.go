//Map declared without a composite literal or make have a value of nil
//Maps can be read when nil, but cannot be written to.
package main

import "fmt"

func main() {
	var soup map[string]int
	fmt.Println(soup == nil)  //Prints true

	measurment, ok := soup["pumpkin"]
	if ok {
		fmt.Println(measurment)
	}
	for ingredient, measurment := range soup {
		fmt.Println(ingredient, measurment)
	}
}