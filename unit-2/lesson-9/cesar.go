//Decipher famous quote from Julius Caesar
package main

import "fmt"

func main() {
    message := "L fdph, L vdz, L frqtxhuhg"

    for i := 0; i < len(message); i++ {
        letter := message[i]
        if letter >= 'a' && letter <= 'z'{ //if any lowercase letter
            letter -= 3  //subtract letter by three
            if letter < 'a'{  //if letter 'a' 
                letter += 26  //wrap back around 
            }
        } else if letter >= 'A' && letter <= 'Z' { //if any uppercase 
            letter -=3  //subtract letter by three
            if letter < 'A' {  //if letter 'A'
                letter += 26  //wrap back around
            }
        }
        fmt.Printf("%c", letter)
    }
}
