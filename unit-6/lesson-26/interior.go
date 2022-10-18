//Introducing interior pointers, they determine the mem. addr. of a field inside a structure.
package main

import "fmt"

type stats struct {
	level 				int
	endurance, health	int
}

type character struct {
	name 	string
	stats 	string
}

func levelUp(s *stats) {  //*stats is the interior pointer
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.health = 5 * s.endurance
}

func main() {
	player := character{name: "Matthias"}
	levelUp(&player.stats)  //Pointer to interior of struct
	fmt.Printf("%+v\n", player.stats)
}