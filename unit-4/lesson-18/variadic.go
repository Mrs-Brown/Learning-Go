//Declaring variadic function
package main

import "fmt"

func terraform(prefix string, worlds ...string) []string {  //Uses ellipsis to declare as variadic
  newWorlds := make([]string, len(worlds))  //Makes a new preallocated slice rather than modifying worlds directly

  for i := range worlds {
    newWorlds[i] = prefix + " " + worlds[i]
  }
  return newWorlds
}

func main() {
  twoWorlds := terraform("New", "Venus", "Mars")
  fmt.Println(twoWorlds)  //Prints [New Venus New Mars]

  planets := []string{"Venus", "Mars", "Jupiter"}
  newPlanets := terraform("New", planets...)
  fmt.Println(newPlanets)
}

