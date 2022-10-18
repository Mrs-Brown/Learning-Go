//If a pointer isn't pointing anywhere, attempting to dereference the pointer will crash the program
package main

import "fmt"

func main() {
	var nowhere *int
	fmt.Println(nowhere)  //Prints <nil>
	//fmt.Println(*nowhere)  //Panic: nil pointer dereference
	if nowhere != nil {
		fmt.Println(*nowhere)  //Instead, dereference with an if statement
	}

}