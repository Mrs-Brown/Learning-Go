//An empty slice and a nil slice aren't equivalent
//but they can often be used interchangeably
package main

import "fmt"

func main() {
	soup := mirepoix(nil)
	fmt.Println(soup)  //Prints [onion carrot celery]
}

func mirepoix(ingredients []string) []string {
	return append(ingredients, "onion", "carrot", "celery")
}