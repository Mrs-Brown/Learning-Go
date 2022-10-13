// Using same structs in a more efficient way
package main

import "fmt"

type report struct {
	sol         int
	temperature //struct of type temperature
	location    //struct of type location
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
	report := report{
		sol:         15,
		temperature: temperature{high: -1.0, low: -78.0},
		location:    location{-4.5895, 137.4417},
	}
	fmt.Printf("average %v celsius\n", report.average())
	fmt.Printf("average %v celsius\n", report.temperature.average())
	fmt.Printf("%v celsius\n", report.high)
	report.high = 32
	fmt.Printf("%v celsius\n", report.temperature.high) //Changes in report.high are reflected in temperature.high
}
