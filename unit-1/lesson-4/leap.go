//Modifying refactor.go to handle leap years and generate a random year instead of 2018
package main

import (
    "fmt"
    "math/rand"
)

var era = "AD"

func main() {
    count := 0

    for count < 10 {

        year := rand.Intn(9999) //Random year from 0-9999
        month := rand.Intn(12) + 1
        daysInMonth := 31
    
        //A leap year is evenly divisible by 4 but not evenly divisible by 100
        //Or any year evenly divisible by 400
        leap := year%400 == 0 || (year%4 == 0 && year%100 != 0)
    
        switch month {
        case 2:
            if leap {
                daysInMonth = 29
            } else {
                daysInMonth = 28
            }
        case 4, 6, 9, 11:
            daysInMonth = 30
        }
        day := rand.Intn(daysInMonth) + 1
        fmt.Println(era, year, month, day)
        count++
    }
}
