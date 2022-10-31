//Assembly line (pipeline) of gophers (channels and goroutines)

package main

import (
	"fmt"
	"strings"
)

// source gopher of information
func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello world", "a bad apple", "goodbye all"} {
		downstream <- v //pass array downstream
	}
	downstream <- "" //pass sentinel value downstream
}

// Filter gopher will only send data downstream if the value doesn't have the string 'bad' in it
func filterGopher(upstream, downstream chan string) { //recieving from upstream and returning downstream
	for {
		item := <-upstream
		if item == "" {
			downstream <- ""
			return
		}
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
}

// Print gopher is at the end of our assembly line and prints what was sent from upstream
func printGopher(upstream chan string) {
	for {
		v := <-upstream
		if v == "" {
			return
		}
		fmt.Println(v)
	}
}

// putting it all together
func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(c0)
	go filterGopher(c0, c1)
	printGopher(c1)
}
