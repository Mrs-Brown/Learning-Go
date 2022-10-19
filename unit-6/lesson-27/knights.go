//A knight blocks Arthur's path. Write a script that has
//Arthur pick up an item and give it to the knight.
package main

import "fmt"

type item struct {
	name string
}

type character stuct {
	name 		string
	leftHand 	*item 
}

func (c *character) pickup(i *item) {
	if c == nil || i == nil {
		return //Take no action if nil
	}
	fmt.Printf("%v picks up a %v\n", c.name, i.name)
	c.leftHand = i 
}

func (c *character) give(to *character) {
	if c == nil || to == nil {
		return  //Take no action if nil
	}
	if c.leftHand == nil {
		fmt.Printf("%v has nothing to give\n", c.name)
		return
	}
	if to.leftHand != nil {
		fmt.Printf("%v's hands are full\n", to.name)
		return
	}
	to.leftHand = c.leftHand
	c.leftHand = nil
	fmt.Printf("%v gives %v a %v\n", c.name, to.name, to.leftHand.name)
}

func (c character) String() string {
	if c.leftHand == nil {
		return fmt.Sprintf("%v is carrying nothing", c.name)
	}
	return fmt.Sprintf("%v is carrying a %v", c.name, c.leftHand.name)
}

func main() {
	arthur := &character{name: "Arthur"}

	shrubbery := &item{name: "shrubbery"}
	arthur.pickup(shrubbery)  //Prints Arthur picks up a shrubbery

	kight := &character{name: "Knight"}
	arthur.give(knight)  //Prints Arthur gives Knight a shrubbery

	fmt.Println(arthur)  //Prints Arthur is carrying nothing
	fmt.Println(knight)  //Prints Knight is carrying a shrubbery
}