//Using composite literal with larger array
package main

func main() {
    planets := [...]string{  //Go compiler counts the elements
        "Mercury",
        "Venus",
        "Earth",
        "Mars",
        "Jupiter",
        "Saturn",
        "Uranus",
        "Neptune",  //Trailing comma is required
    }
}
