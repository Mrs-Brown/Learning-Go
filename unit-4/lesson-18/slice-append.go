// Dump slice length, capacity, and contents
package main

import (
	"fmt"
)

// Dump slice length, capacity, and contents
func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}

func main() {
	dwarfs1 := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dump("dwarfs1", dwarfs1) //Prints dwarfs: length 5, capacity 5 [contents]

	dwarfs2 := append(dwarfs1, "Orcus")
	dump("dwarfs2", dwarfs2)

	dwarfs3 := append(dwarfs2, "Salacia", "Quaoar", "Sedna")
	dump("dwarfs3", dwarfs3)
}
