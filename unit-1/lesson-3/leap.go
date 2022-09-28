//Calucluate whether a year is a leap year or not with if statements
//and &&, || statements
package main

import "fmt"

func main() {
    fmt.Println("The year is 2100, should you leap?")

    var year = 2100
    //A leap year is evenly divisible by 4 but not evenly divisible by 100
    //Or any year evenly divisible by 400
    var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)

    if leap {
        fmt.Println("Look before you leap!")
    } else {
        fmt.Println("Keep your feet on the ground")
    }
}
