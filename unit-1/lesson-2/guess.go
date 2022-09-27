//Print random number from 1 through 10 using rand package 
package main

import (
    "fmt"
    "math/rand"
)

func main() {
    var num = rand.Intn(10) + 1 //random val 1 thru 10 (not 0 thru 9)
    fmt.Println(num)
}
