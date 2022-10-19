//Can use an if statement to handle errors

package main

import (
	 "fmt"
	 "os"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, file := range files {
		fmt.Println(file.Name())  //Prints a list of directories
	}
}