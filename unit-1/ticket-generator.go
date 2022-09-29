//Building a "ticket to Mars" ticketing system
//Format:
//Spaceline             Days    Trip type   Price
//-----------------------------------------------
//Generating 10 tickets
package main

import (
    "fmt"
    "math/rand"
)

const secondsPerDay = 86400 //Constant amount of seconds in a day

func main() {
    
    //format display for tickets
    fmt.Println("Spaceline         Days  Trip type Price")
    fmt.Println("---------------------------------------------")
   
    const distance = 118346487 //Distance from Earth to Mars on 9/29/22
    spaceline := "" //Placeholder for spaceline company
    trip := ""  //Placeholder for trip type
    price := 0  //Placeholder for trip price

    for count := 0; count < 10; count++ { //Print 10 tickets
        
        switch rand.Intn(3) { //Choose from 1 of 3 spaceline companies
        case 0:
            spaceline = "Space Adventures"
        case 1:
            spaceline = "SpaceY"
        case 2:
            spaceline = "NASA"
        }

        speed := rand.Intn(15) + 16 //Random int from 16-30 km/s
        days := distance / speed / secondsPerDay //days
       
        if speed < 23 { //Calculate price based off speed
         price = rand.Intn(22) + 36 //Random int from 36 to 44 million
        } else if speed > 24 {
         price = rand.Intn(25) + 44 //Random int from 44 to 50 million
        } else {
         price = 20 + speed
        }


        if rand.Intn(2) == 1 { //Calculate whether trip is round trip or not
            trip = "Round-trip"
            price *= 2
        } else {
            trip = "One-way"
        }

        fmt.Printf("%-16v %4v %-10v %4v\n", spaceline, days, trip, price)

    }

}
