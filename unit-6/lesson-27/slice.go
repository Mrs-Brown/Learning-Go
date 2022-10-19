//range, len, append work with nil slices
package main

import "fmt"

func main() {
	var soup []string
	fmt.Println(soup == nil)  //Prints true

	for _, ingredient := range soup {
		fmt.Println(ingredient)
	}

	fmt.Println(len(soup))  //Prints 0

	soup = append(soup, "pumpkin", "butter", "milk")
	fmt.Println(soup)  //Prints [pumpkin butter milk]
}