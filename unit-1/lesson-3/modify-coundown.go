//Modified countdown.go to interupt randomly
package main

import(
    "fmt"
    "time"
    "math/rand"
)

func main() {
    var count = 10

    for count > 0 {
        fmt.Println(count)
        time.Sleep(time.Second)
        if rand.Intn(100) + 1 == 0{  //A 1 in 100 chance (1-100)
            break                    //that the loop will break
        }
        count--
    }
    if count == 0 {
        fmt.Println("Liftoff!")
    } else {
        fmt.Println("Launch failed")
    }
}
