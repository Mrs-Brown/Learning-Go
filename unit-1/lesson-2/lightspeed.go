//Calculate time to get to mars using constants and variables
package main

import "fmt"

func main() {
    const lightSpeed = 299792 //km per second
    var distance = 56000000 //km

    fmt.Println(distance/lightSpeed, "seconds") //Prints 186 seconds

    distance = 401000000 //km

    fmt.Println(distance/lightSpeed, "seconds") //Prints 1337 seconds
}
