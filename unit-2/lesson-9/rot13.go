//Decoding a message in rot13 format
package main

import "fmt"

func main() {
    message := "uv vagreangvbany fcnpr fgngvba"
    
    for i := 0; i < len(message); i++ { //Iterate through each char
        letter := message[i]
        if letter >= 'a' && letter <= 'z' {
            letter = letter + 13
            if letter > 'z' {
                letter = letter - 26
            }
        }
        fmt.Printf("%c", letter)  
    }

}

