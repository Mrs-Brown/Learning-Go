//Now trying a method using pointer recievers
package main

import "fmt"

type laser int

type talker interface {
	talk() string
}

func (l *laser) talk() string {
	return strings.Repeat("pew ", int(*1))
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func main() {
	pew := laser(2)
	shout(&pew)  //Prints PEW PEW
}