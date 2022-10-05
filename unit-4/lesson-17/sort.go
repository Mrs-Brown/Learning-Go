//Sorting a slice by alphabetizing 
package main

import (
    "fmt"
    "sort"
)

func main() {
    planets := []string{  //Init slice
        "Mercury", "Venus", "Earth", "Mars",
        "Jupiter", "Saturn", "Uranus", "Neptune",
    }

    sort.StringSlice(planets).Sort()  //Sort planets alphabetically
    fmt.Println(planets)
}
