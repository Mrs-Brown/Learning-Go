// Using make to preallocate slices
package main

func main() {
	dwarfs := make([]string, 0, 10) //Make a slice of 0 len and 10 cap
	dwarfs = append(dwarfs, "Ceres", "Pluto", "Haumea", "Makemake", "Eris")
	//Preallocating with make can set an initial capacity, therefore avoiding additional allocations
	//and copies to enlarge the underlying array
}
