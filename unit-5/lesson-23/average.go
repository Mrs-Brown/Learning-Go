// Using same structs to calculate average temperature
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

// Calculate average temperature
func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

// method that forwards to the real implementation
func (r report) average() celsius {
	return r.temperature.average()
}

func main() {
	bradbury := location{-4.5895, 137.4417}
	t := temperature{high: -1.0, low: -78.0}
	report := report{sol: 15, temperature: t, location: bradbury}

	fmt.Printf("%+v\n", report)
	fmt.Printf("a balmy %v celsius\n", report.temperature.high)
	fmt.Printf("average %v celsius\n", t.average())
	fmt.Printf("Same data again with method forwarding: %v\n", report.average())
}
