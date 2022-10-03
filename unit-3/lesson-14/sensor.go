//Assign functions to variables
package main

import (
    "fmt"
    "math/rand"
)

type kelvin float64

func fakeSensor() kelvin {
    return kelvin(rand.Intn(151) +150) //Pseudorandom var for fake sensor results
}

func realSensor() kelvin {
    return 0 //Will implement once sensor is obtained for Mars
}

func main() {
    sensor := fakeSensor //Assign function to variable
    fmt.Println(sensor())

    sensor = realSensor //Assign function to same variable
    fmt.Println(sensor())
}

//Sensor variable is of type function
