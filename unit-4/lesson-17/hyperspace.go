//Create a function that removes the space between planets for instantaneous travel
package main

import (
    "fmt"
    "strings"
)

//hyperspace removes the space surrounding worlds
func hyperspace(worlds []string) {  //Argument is a slice, not array
    for i := range worlds {
        worlds[i] = strings.TrimSpace(worlds[i])  //Invoke TrimSpace method
    }
}

func main() {
    planets := []string{"  Venus    ", "Earth  ", " Mars"}  //Slice where planets are surrounded by space
    hyperspace(planets)  

    fmt.Println(strings.Join(planets, ""))  //Join the elements to print VenusEarthMars
}
