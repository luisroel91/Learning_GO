package main

import (
	"fmt"
)

//Helper function to print arrays
func printArray(a [3][2]string) {
	for _, column := range a {
		for _, row := range column {
			fmt.Printf("%s ", row)
		}
		fmt.Printf("\n")
	}
}

// Multidimensional (list of lists) arrays can be declared
// using the following format:
//
// array := [number of rows][number of columns] data type {
// 		{value, value},
// 		{value, value},
//      {value, value}, <- MUST have trailing comma
//     } <- ending bracket

func main() {
	// Declare array with two columns, three rows of type string
	a := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		// Need trailing comma due to lexer rules
		{"pigeon", "peacock"},
	} // Ending bracket
	printArray(a)

	// We can also define the array without values, later
	// inserting them by addressing them using their
	// column and row index numbers
	//
	// For example, if we wanted to reassign the value in the first
	// column second row we would do:
	//
	// array[1][0] = some value

	var b [3][2]string
	b[0][0] = "apple"
	b[0][1] = "samsung"
	b[1][0] = "microsoft"
	b[1][1] = "google"
	b[2][0] = "AT&T"
	b[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printArray(b)
}
