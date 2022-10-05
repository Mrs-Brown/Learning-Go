//If you begin by initializing a slice, Go in the background will declare an array and make a slice that views all of its elements
package main

import "fmt"

func main() {
    dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
    fmt.Println(dwarfs)
}
