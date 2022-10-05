// Using maps
package main

import "fmt"

func main() {
	temperature := map[string]int{ //map; key type: string; value type: int
		"Earth": 15,  //Composite literals are key-value pairs for maps
		"Mars":  -65, //Trailing comma necessary
	}

	temp := temperature["Earth"]
	fmt.Printf("On average the Earth is %v °C.\n", temp)

	temperature["Earth"] = 16
	temperature["Venus"] = 464

	fmt.Println(temperature)
}
