//Vigenere cipher uses a keyword to encode/decode a message
//Here we're encoding a message
package main

import "fmt"

func main() {
    plainText := "WEDIGYOULUVGO"
    keyword := "GOLANG"
    cipherText := ""
    keyIndex := 0

    for i := 0; i < len(plainText); i++ {
    //A=0, B=1...Z=25
    c := plainText[i] 
    if c >= 'A' && c <= 'Z'{
        c -= 'A'
        k := keyword[keyIndex] - 'A'

        //cipher letter + key letter
        c = (c+k)%26 + 'A' //modulus to wrap around

        keyIndex++
        keyIndex %= len(keyword)
    }
    
    cipherText += string(c)
    }
    fmt.Println(cipherText)
}
