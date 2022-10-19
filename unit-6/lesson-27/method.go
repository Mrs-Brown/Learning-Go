//Whether a pointer is dereferenced explicitly or implicity
//by accessing a field of the struct, a nil value will panic
package main

import "fmt"

type person struct {
	age int
}

func (p *person) birthday() {
	//p.age++  //nil pointer dereference, program crashes
	if p == nil {
		return  //If nil return, program runs
	}
	p.age++
}

func main() {
	var nobody *person
	fmt.Println(nobody)  //Prints <nil>

	nobody.birthday()
}