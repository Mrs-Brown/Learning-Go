// Using structs for a weather report
package main

import "fmt"

type report struct {
	sol         int
	temperature temperature //struct of type temperature
	location    location    //struct of type location
}

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

type celsius float64

func main() {
	bradbury := location{-4.5895, 137.4417}
	t := temperature{high: -1.0, low: -78.0}
	report := report{sol: 15, temperature: t, location: bradbury}

	fmt.Printf("%+v\n", report)
	fmt.Printf("a balmy %v celsius\n", report.temperature.high)
}
