//Capstone priject to create a guessing game for 1-10
package main

import(
    "fmt"
    "math/rand"
)

func main() {
    var correct = 5 //The correct number to guess 

    fmt.Println("Let's play a guessing game.")
    fmt.Println("I don't know what the correct value is, but you will.")
    fmt.Printf("The correct value is: %v\n", correct)
    fmt.Println("I will guess until I'm correct, okay?")

    for {
        var guess = rand.Intn(10) + 1 //Guess a random value from 1-10
        if guess < correct {
            fmt.Printf("%v is too small.\n", guess)
        } else if guess > correct {
            fmt.Printf("%v is too big.\n", guess)
        } else {
            fmt.Printf("You got it! %v\n", guess)
            break
        }    
    }
}

