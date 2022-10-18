//Func martian and pointer to martian satisfy the type talker interface
package main

import "fmt"

type talker interface {
	talk() string
}

type martian struct{}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func (m martian) talk() string {
	return "nack nack"
}

func main() {
	shout(martian{})  //Prints NACK NACK
	shout(&martian{})  //Prints NACK NACK
}