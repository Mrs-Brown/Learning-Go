//Orbit a 360 degree circle and stop randomly
package main

import (
    "fmt"
    "math/rand"
)


func main() {
    var degrees = 0

    for {
        fmt.Println(degrees)

        degrees++
        if degrees >= 360 {  //If degrees goes over 360
            degrees = 0      //Revert back to zero
            if rand.Intn(2) == 0 {   //If a random value form 0 to 1 equals 0
                break               //break out of the loop
            }
        }
    }
}
