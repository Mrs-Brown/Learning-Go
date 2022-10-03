//The calibrate function accepts a fake or real sensor as a paramerter and returns a replacement function
package main

import "fmt"

type kelvin float64
//sensor function type
type sensor func() kelvin

func realSensor() kelvin {
    return 0 //Will implement sensor
}

func calibrate(s sensor, offset kelvin) sensor {
    return func() kelvin {  //Declares and returns anon function
        return s() + offset
    }
}

func main() {
    sensor := calibrate(realSensor, 5)
    fmt.Println(sensor())  //Prints 5
}

