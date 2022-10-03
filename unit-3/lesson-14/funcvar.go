//Scope for assigning anon func to var
package main

import "fmt"

func main() {
    f := func(message string) { //Assigns anon function to var
        fmt.Println(message)
    }
    f("Go to the party") //Calling function through variable
}
