//Example usage of rune
package main

import "fmt"

func main() {
    //var pi rune = 960
    //var alpha rune = 940
    //var omega rune = 969
    var name rune = 233
    var smile rune = 128515
    var bang byte = 33

    //fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang)
    //fmt.Printf("%c%c%c%c\n", pi, alpha, omega, bang)
    fmt.Printf("T%ca%c%c\n", name, bang, smile)
}
