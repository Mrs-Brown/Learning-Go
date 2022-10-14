// Print who is the current admin of Nasa
package main

import "fmt"

func main() {
	var administrator *string

	scolese := "Christopher J. Scolese"
	administrator = &scolese
	fmt.Println(*administrator) //Prints Christopher J. Scolese

	bolden := "Charles F. Bolden"
	administrator = &bolden
	fmt.Println(*administrator) //Prints Charles F. Bolden
}
