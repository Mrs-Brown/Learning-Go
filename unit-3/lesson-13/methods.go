//Using methods
package main

import "fmt"

type celsius float64
type fahrenheit float64
type kelvin float64

func (c celsius) fahrenheit() fahrenheit { //Method named fahrenheit that takes a celsius type var and returns fahrenheit type var
    return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (c celsius) kelvin() kelvin {
    return kelvin(c + 273.15)
}

func (f fahrenheit) celsius() celsius {
    return celsius((f - 32.0) * 5.0 / 9.0)
}

func (f fahrenheit) kelvin() kelvin {
    return f.celsius().kelvin()
}

func (k kelvin) celsius() celsius {
    return celsius(k - 273.15)
}

func (k kelvin) fahrenheit() fahrenheit {
    return k.celsius().fahrenheit()
}

func main() {
    var k kelvin = 294.0
    c := k.celsius()  //Calling celsius method that takes a kelvin var and returns celsuis var
    fmt.Printf("%6.2f°K is %6.2f°C\n", k, c)
}
