// Introducing channels
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) //Makes the channel to communicate over
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}
	for i := 0; i < 5; i++ {
		gopherID := <-c //Recieves a value from a channel
		fmt.Println("gopher ", gopherID, " has finished sleeping")
	}
}

func sleepyGopher(id int, c chan int) { //Declares the channel as an argument
	time.Sleep(3 * time.Second)
	fmt.Println("...", id, " snore ...")
	c <- id //Sends a value back to main
}
