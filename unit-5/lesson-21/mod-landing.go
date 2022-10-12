// Display three rover landing sites in {JSON
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type location struct { //Structs with JSON tags
		Name string  `json:"name"`
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}

	locations := []location{ //Slice of structs
		{Name: "Bradbury Landing", Lat: -4.5895, Long: 137.4417},
		{Name: "Columbia Memorial Station", Lat: -14.5684, Long: 175.472636},
		{Name: "Challenger Memorial Station", Lat: -19462, Long: 354.4734}, //Note trailing comma
	}

	bytes, err := json.MarshalIndent(locations, "", "  ") //Format JSON output
	if err != nil {                                       //If there is an error...
		fmt.Println(err) //Print the error...
		os.Exit(1)       //and exit
	}

	fmt.Println(string(bytes)) //Output JSON as strings and print
}
