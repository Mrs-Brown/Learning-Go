//A spaceship lacks a warp drive, but it will coast to Mars at a 
//respectable 100,800 km/h. An ambitious launch date of Jan. 2025 
//would place Mars and Earth 96,300,000 km apart. How many days 
//would it take to reach Mars? Modify lightspeed.go to find out

package main

import "fmt"

func main() {
    const speed = 100800 //km/h 
    const days = 24 //hours per day
    var distance = 96300000 //km

    fmt.Println(distance/speed/days, "days") //Prints 39 days

}
