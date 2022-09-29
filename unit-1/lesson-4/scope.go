//Example of scope using a for loop
package main

import (
    "fmt"
    "math/rand"
)

func main() {
    var count = 0

    for count < 10 {  //Scope begins
        var num = rand.Intn(10) + 1
        fmt.Println(num)
        count++
    }  //Scope ends
}
