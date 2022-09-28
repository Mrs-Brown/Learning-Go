//Using for loop for a countdown timer
package main

import (
    "fmt"
    "time"
)

func main() {
    var count = 10  //Declare and initialize

    for count > 0 { //A condition
        fmt.Println(count)
        time.Sleep(time.Second)
        count--     //Decrements count; otherwise infinite loop error
    }
    fmt.Println("Liftoff!") 
}
