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

	//comma, ok syntax is used to distinguish between a key not existing in the map
	//versus being present in the map
	moon := temperature["Moon"] //Assigning key, but no value
	fmt.Println(moon)           //Prints 0 because no value is assigned

	if moon, ok := temperature["Moon"]; ok { //The comma, ok syntax
		fmt.Printf("On average the moon is %v°C.\n", moon)
	} else {
		fmt.Println("Where is the moon?")
	}
}
