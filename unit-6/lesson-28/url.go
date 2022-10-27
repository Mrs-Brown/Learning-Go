//Lesson capstone
package main

import (
	"fmt"
	"os"
)

func main() {
	u, err := url.Parse("https://a b.com")

	if err != nil {
		fmt.Print(err)  //Prints parse https://a b.com/:invalid character " " in host name
		fmt.Printf("%#v\n", err)  //Prints &url.Error{Op:"parse", URL:"https://a b.com/", Err:" "}
		
		if e, ok := err.(*url.Error); ok {  
			fmt.Println("Op:", e.Op)  //Prints Op: parse
			fmt.Println("URL:", e.URL)  //Prints URL: https://a b.com/
			fmt,Println("Err:", e.Err)  //Prints Err: invalid character " " in host name
		}
		os.Exit(1)
	}
	fmt.Println(u)  
}