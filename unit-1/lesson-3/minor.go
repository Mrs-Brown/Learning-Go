//Compare whether or not age is a minor
package main

import "fmt"

func main() {
    fmt.Println("There is a sign near the entrance that reads 'No Minors'.")

    var age = 41
    var minor = age < 18 //A minor is less than 18 

    fmt.Printf("At age %v, am I a minor? %v\n", age, minor)
}
