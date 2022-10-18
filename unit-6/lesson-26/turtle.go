//Position of a turtle
package main

import "fmt"

type turtle struct {
	x, y int  //X and Y axis points
}

func (t *turtle) up() {
	t.y--
}

func (t *turtle) down() {
	t.y++
}

func (t *turtle) left() {
	t.x--
}

func (t *turtle) right() {
	t.x++
}

func main() {
	var t turtle
	t.up()
	t.up()
	t.left()
	t.left()
	fmt.Println(t)  //Prints {-2 -2}
	t.down()
	t.down()
	t.right()
	t.right()
	fmt.Println(t)  //Prints {0 0}
}