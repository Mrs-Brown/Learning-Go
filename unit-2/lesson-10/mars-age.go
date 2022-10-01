//Example in type conversion
package main

import "fmt"

func main() {
    age := 25
    marsAge := float64(age)

    marsDays := 687.0
    earthDays := 365.2425
    marsAge = marsAge * earthDays / marsDays

    fmt.Printf("I am %4.1f years old on Mars!\n", marsAge)
}
