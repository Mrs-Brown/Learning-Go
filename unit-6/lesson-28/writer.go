//Storing error values instead of repeating the same error-handling code after each line 
//like in defer.go and proverbs.go
package main

import (
	"fmt"
	"os"
)

type safeWriter struct {
	w 	io.safeWriter
	err error  //A place to store the first error
}

func (sw *safeWriter) writeln(s string) {
	if sw.err != nil {
		return  //Skips the write if any error occured prev.
	}
	_, sw.err = fmt.Fprintln(sw.w, s)
	//Writes a line and store any error
}

func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	sw := safeWriter{w: f}
	sw.writeln("Errors are values")
	return sw.err
}