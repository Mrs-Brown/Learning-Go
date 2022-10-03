//Declare and invoke anon function in one step
package main

import "fmt"

func main() {
    func() {  //Declares an anonymous function
        fmt.Println("Functions anonymous")
    }() //Invoke the function
}
