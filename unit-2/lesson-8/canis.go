//Use constants to convert the distance of Canis Major Dwarf to the sun
//in light years
package main

import "fmt"

func main() {
    const canis = 236000000000000000
    const lightSpeed = 299792
    const secondsPerDay = 86400
    const daysPerYear = 365

    const years = canis / lightSpeed / secondsPerDay / daysPerYear

    fmt.Println("Canis Major Dwarf is", years, "light years away!" )
}
