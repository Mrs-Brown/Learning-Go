//Cipher spanish message in rot13
package main

import "fmt"

func main() {
    message := "Hola Estación Espacial Internacional"
    fmt.Println("The secret message we need to send to the aliens is:")
    fmt.Println(message)
    fmt.Println("The encoded message is:")

    for _, c:= range message {
        if c >= 'a' && c <= 'z' { //If a lowercase letter
            c += 13  //Add 13
            if c > 'z' { //if z
                c = c - 26 //wrap back around
            }
        } else if c >= 'A' && c <= 'Z' { //if uppercase letter
            c += 13 //Add 13
            if c > 'Z' {  //if Z
                c = c - 26  //wrap back around
            }
        } 
    fmt.Printf("%c", c)
    }
}
