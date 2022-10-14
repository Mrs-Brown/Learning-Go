package main

import "fmt"

func main() {
	canada := "Canada"

	var home *string
	fmt.Printf("home is a %T\n", home) //pointer type *string

	home = &canada
	fmt.Println(*home) //prints Canada
}
