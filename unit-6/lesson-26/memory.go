package main

import "fmt"

func main() {
	answer := 42
	fmt.Println(&answer) //Print memory address of var

	address := &answer    //Give address value of memory address location
	fmt.Println(*address) //Points to value at memory address
}
