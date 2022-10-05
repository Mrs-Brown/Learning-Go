//Combiining lesson into program
package main

import "fmt"

//Planets attaches methods to []string slice
type Planets []string

func (planets Planets) terraform() {  //Method that recieves type Planet 
    for i := range planets {  //For each element
        planets[i] = "New " + planets[i]  //Add word "New" to planet
    }
}

func main() {
    planets := []string{  //Init slice
        "Mercury", "Venus", "Earth", "Mars",
        "Jupiter", "Saturn", "Uranus", "Neptune",
    }
    Planets(planets[3:4]).terraform()  //Invoke method
    Planets(planets[6:]).terraform()  //Invoke method
    fmt.Println(planets)
}
