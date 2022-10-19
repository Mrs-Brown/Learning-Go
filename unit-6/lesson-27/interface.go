//When a var is declared to be an interface type
//without an assignment, the zero value is nil
package main

import "fmt"

func main() {
	var v interface{}
	fmt.Printf("%T %v %v\n", v, v, v == nil)  //Prints <nil> <nil> true
	//Both the interface type and value are nil

	var p *int
	v = p
	fmt.Printf("%T %v %v\n", v, v, v == nil)  //Prints *int <nil> false

	fmt.Printf("%#v\n", v)  //Prints(*int)(nil)
    //%#v is shorthand to see both type and value
	//To avoid surprises when comparng interfaces to nil, use the nil 
	//identifier explicity, rather than pointing to a var that contains nil
	
}
