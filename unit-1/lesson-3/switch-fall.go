//Switch fallthrough expression example
package main

import "fmt"

func main() {
    var room = "lake"

    switch {   //Expressions are in each case
    case room == "cave":
        fmt.Println("You find yourself in a dimly lit cave.")
    case room == "lake":
        fmt.Println("The ice on the lake seems solid enough.")
        fallthrough  //Fallsthrough to the next case
    case room == "underwater":
        fmt.Println("You find yourself underwater in freezing temperatures.")
    }
}
