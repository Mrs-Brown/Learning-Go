//The previous code in scope-rules was 'smelly' because it had code duplication.
//So let's refactor that. Refactoring is modifying the code without modifying the behavior
package main

import (
    "fmt"
    "math/rand"
)

var era = "AD"

func main() {
    year := 2018
    month := rand.Intn(12) + 1
    daysInMonth := 31

    switch month {
    case 2:
        daysInMonth = 28
    case 4, 6, 9, 11:
        daysInMonth = 30
    }

    day := rand.Intn(daysInMonth) + 1
    fmt.Println(era, year, month, day)
}
