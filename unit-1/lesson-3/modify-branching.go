//Modify if statements to describe beginning of three different adventures
//Cave, entrance, and mountain 
package main

import "fmt"

func main() {
    var room = "cave"

    if room == "cave" {
        fmt.Println("You find yourself in a dimly lit cave.")
    } else if room == "entrance" {
        fmt.Println("You find yourself at the entrance of a cave with a cat blocking your path.")
    } else if room == "mountain" {
        fmt.Println("You notice a village by the ocean down the mountain your on.")
    } else {
        fmt.Println("All around you is white.")
    }
}
