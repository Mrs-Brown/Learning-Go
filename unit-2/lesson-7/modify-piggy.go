//Modify lesson-6 piggy.go to use integers to track the number of cents
//rather than dollars.
package main

import (
    "fmt"
    "math/rand"
)

func main() {
    nickles := 5
    dimes := 10
    quarters := 25
    bank := 0

    for bank < 2000 { //Bank needs to have at least $20
        switch rand.Intn(3){
        case 0:
            bank = bank + nickles
        case 1:
            bank = bank + dimes
        case 2:
            bank = bank + quarters
        }
        
        dollars := bank / 100
        cents := bank % 100
        fmt.Printf("$%d.%02d\n", dollars, cents)
    }
}
