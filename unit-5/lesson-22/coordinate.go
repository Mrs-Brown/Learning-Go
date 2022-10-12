//Attaching methods to structures

package main

import "fmt"

func main() {
	//Bradbury Landing coordinate
	lat := coordinate{4, 35, 22.2, 'S'}
	long := coordinate{137, 26, 30.12, 'E'}
	fmt.Println(lat.decimal(), long.decimal())
}

//coordinate in degrees, min, sec in a N/S/E/W hemisphere
type coordinate struct {
	d, m, s float64
	h       rune
}

// decimal converts a d/m/s coordinate to decimal in degrees
func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}
