//using switch statement
package main

import "fmt"

func main() {
    fmt.Println("There is a cavern entrance here and a path to the east.")
    var command = "go inside"

    switch command {  //Compare cases to command
    case "go east":
        fmt.Println("You head further up the mountain.")
    case "enter cave", "go inside":   //A comma seperated list of possible values
        fmt.Println("You find yourself in a dimly lit cave.")
    case "read sign":
        fmt.Println("The sign reads 'No Minors'.")
    default:
        fmt.Println("Didn't quite get that.")
    }
}
