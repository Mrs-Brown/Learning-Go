// Discover interfaces
package main

import (
	"fmt"
	"time"
)

type stardater interface {
	YearDay() int
	Hour() int
}

type sol int

// stardate returns a fictional measure of time for a given date
func stardate(t stardater) float64 {
	day := float64(t.YearDay())
	h := float64(t.Hour() / 24.0)
	return 1000 + day + h
}

func (s sol) YearDay() int {
	return int(s % 668)
}

func (s sol) Hour() int {
	return 0
}

func main() {
	day := time.Date(2012, 8, 6, 5, 17, 0, 0, time.UTC)
	fmt.Printf("%.1f Curiosity has landed\n", stardate(day))

	s := sol(1422)
	fmt.Printf("%.1f Happy Birthday\n", stardate(s))
}
