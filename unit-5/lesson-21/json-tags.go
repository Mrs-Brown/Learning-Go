package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type location struct { //Structs of json tags
		Lat  float64 `json:"latitude"` //Fields begin with uppercase
		Long float64 `json:"longitude"`
	}

	curisoity := location{-4.5895, 137.4417} //Lat and Long of curiosity rover

	bytes, err := json.Marshal(curisoity)
	exitOnError(err)

	fmt.Println(string(bytes))
}

// exitOnError prints any errors and exits
func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
