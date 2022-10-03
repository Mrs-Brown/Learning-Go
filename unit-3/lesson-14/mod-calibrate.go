//
package main

import (
    "fmt"
    "math/rand"
)

type kelvin float64
type sensor func() kelvin

func fakeSensor() kelvin { //Function that returns kelvin
    return kelvin(rand.Intn(151) + 150)
}

func calibrate(s sensor, offset kelvin) sensor { //Function that includes function as parameter and returns function
    return func() kelvin {  //Return anon function
        return s() + offset //Return invoked function plus var
    }
}

func main() {
    var offset kelvin = 5
    sensor := calibrate(fakeSensor, offset) //Var is result of calibrate function

    for count := 0; count < 10; count++ {
        fmt.Println(sensor())  //Print var sensor
    }
}
