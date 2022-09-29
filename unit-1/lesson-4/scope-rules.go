//further explaining rules of variable scope
package main

import (
    "fmt"
    "math/rand"
)

var era = "AD" //Era is a global variable available throught the package

func main() {
    year := 2018  //short declaration; year and era are in scope

    switch month := rand.Intn(12) + 1; month { //era, year, month in scope
    case 2:
        day := rand.Intn(28) + 1 //era, year, month, day in scope
        fmt.Println(era, year, month, day)
    case 4, 6, 9, 11:
        day := rand.Intn(30) + 1 //New day variable
        fmt.Println(era, year, month, day)
    default:
        day := rand.Intn(31) + 1 //New day variable
        fmt.Println(era, year, month, day)
    } //month and day are out of scope
} //year out of scope
