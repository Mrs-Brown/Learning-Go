//Assign an anonymous function to a var and then use the var like any other function
package main

import "fmt"

var f = func() { //Assigns an anon function to a var
    fmt.Println("Dress up for the masquerade.")
}

func main() {
    f() //Use the variable to call the anon function
}
