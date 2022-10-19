//Pointers are intended for pointing, so using a pointer
//just to provide a nil value isn't the best option.
//One alternative is to declare a small struct with a few methods
package main

import "fmt"

type number struct {
	value int
	valid bool
}

func newNumber(v int) number{
	return number{value: v, valid: true}
}

func (n number) String() string {
	if !n.valid {
		return "not set"
	}
	return fmt.Sprintf("%d", n value)
}

func main() {
	n := newNumber(42)
	fmt.Println(n)  //Prints 42

	e := number{}
	fmt.Println(e)  //Prints not set
}