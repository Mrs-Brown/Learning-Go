//Write a program that randomly places nickles, dimes, and quarters into 
//an empty piggy bank until it contains at least $20. Display the running
//balance of the piggy bank after each deposit formatting it with an
//appropriate width and percision
package main

import (
    "fmt"
    "math/rand"
)

func main() {

    nickle := 0.05
    dime := 0.10
    quarter := 0.25
    bank := 0.0
    
    for bank < 20.00{ //minimum amount of money is $20
         fill := rand.Intn(3) //Random value to add to piggy bank
         
         switch fill{
         case 0:
             bank = bank + nickle
         case 1:
             bank = bank + dime
         case 2:
             bank = bank + quarter
         }
        fmt.Printf("%5.2f\n",bank)
    }
}
