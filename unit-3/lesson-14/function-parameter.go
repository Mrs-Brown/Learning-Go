//Passing functions to other functions
package main

import (
    "fmt"
    "math/rand"
    "time"
)

type kelvin float64

func measureTemperature(samples int, sensor func() kelvin) { //Accepts function as second parameter
    for i := 0; i < samples; i++ {
        k := sensor()
        fmt.Printf("%v °K\n", k)
        time.Sleep(time.Second)
    }
}

func fakeSensor() kelvin {
    return kelvin(rand.Intn(151) + 150)
}

func main() {
    measureTemperature(3, fakeSensor) //Passes fakeSensor function to measureTemperature
}
