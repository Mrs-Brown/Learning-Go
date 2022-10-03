//Temperature table for conversion between celsius to fahrenheit and vice versa
package main

import "fmt"

type celsius float64
type fahrenheit float64
type getRowFn func(row int) (string, string) //Type function that recieves int and returns string

const(
    line = "===================="
    rowFormat = "| %8s | %8s |\n"
    numberFormat = "%.1f"
)


func (c celsius) fahrenheit() fahrenheit {  //Method that recieves celcius and returns fahrenheit
    return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (f fahrenheit) celsius() celsius {
    return celsius((f - 32.0) * 5.0 / 9.0)  //Method that recieves fahrenheit and returns celsius
}

//drawTable draws a two column table
func drawTable(hdr1, hdr2 string, rows int, getRow getRowFn) {
    fmt.Println(line)
    fmt.Printf(rowFormat, hdr1, hdr2)
    fmt.Println(line)
    for row := 0; row < rows; row++ {
        cell1, cell2 := getRow(row) //Assign cells
        fmt.Printf(rowFormat, cell1, cell2) //Print cells in rowFormat
    }
    fmt.Println(line)
}

//Celsius to fahrenheit function that returns formated cells
func ctof(row int) (string, string) {
    c := celsius(row*5 - 40)  //Cell format
    f := c.fahrenheit()  //Calling method
    cell1 := fmt.Sprintf(numberFormat, c)
    cell2 := fmt.Sprintf(numberFormat, f)
    return cell1, cell2
}

//Fahrenheit to celsius function that returns formated cells
func ftoc(row int) (string, string) {
    f := fahrenheit(row*5 - 40)  //Cell format
    c := f.celsius()  //Calling method
    cell1 := fmt.Sprintf(numberFormat, c)
    cell2 := fmt.Sprintf(numberFormat, f)
    return cell1, cell2
}

func main() {
    drawTable("C", "F", 29, ctof)
    fmt.Println()
    drawTable("F", "C", 29, ftoc)
}
