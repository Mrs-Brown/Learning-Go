//Malacandra is another name for Mars. Write a program to determine how
//fast a ship would need to travel (km/h) in order to reach Malacandra
// in 28 days. Assume a distance of 56,000,000 km.
package main

import "fmt"

func main() {
    const distance, days, hours = 56000000, 28, 24
    var speed, total = 0, 0
    
    total = days * hours //Multiply hours by days for total hours

    speed = distance * total //Physics formula s=d/t 

    fmt.Printf("It will take %v hours", total)
    fmt.Printf(" to cover a distance of %v km ", distance)
    fmt.Printf("over a constant speed of %v for the star ship to reach Malacandra.", speed)

}
