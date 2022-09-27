//Branching with if statements
package main

import "fmt"

func main() {
    var command = "go east"

    if command  == "go east" {
        fmt.Println("You head further up the mountain.")
    } else if command == "go inside" {
        fmt.Println("You enter the cave where you live the rest of your life.")
    } else {
        fmt.Println("Need more information.")
    }
}
