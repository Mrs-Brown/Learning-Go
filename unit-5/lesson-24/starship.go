// more Polymorphism with interfaces
package main

import (
	"fmt"
	"strings"
)

var t interface {
	talk() string
}

type martian struct{}

type laser int

type talker interface {
	talk() string
}

type starship struct {
	laser
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func (m martian) talk() string {
	return "nack nack"
}

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

func main() {
	s := starship{laser(3)}

	fmt.Println(s.talk()) //prints pew pew pew
	shout(s)              //prints PEW PEW PEW
}
