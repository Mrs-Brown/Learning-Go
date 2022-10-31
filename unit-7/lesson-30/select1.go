// Introducing select statement
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int) //Make channel to communicate over
	timeout := time.After(2 * time.Second)
	for i := 0; i < 5; i++ {
		select {
		case gopherID := <-c: //Waits for a gopher to wake up
			fmt.Println("gopher ", gopherID, " has finished sleeping")
		case <-timeout: //Waits for time to run out
			fmt.Println("my patience ran out")
			return //Gives up and returns
		}
	}
} //Our patience runs out before any gophers wake up

// sleepyGOpher function will allow some gophers to wake up
func sleepyGopher(id int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	c <- id
}
