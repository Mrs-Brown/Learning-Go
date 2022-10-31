// Introducing goroutine
package main

import (
	"fmt"
	"time"
)

func main() {
	go sleepyGopher()           //The goroutine is started
	time.Sleep(4 * time.Second) //Waiting for the gopher to snore
} //All goroutines are stopped here

func sleepyGopher() {
	time.Sleep(3 * time.Second) //The gopher sleeps
	fmt.Println("...snore...")
}
