package main

import "fmt"

func main() {

	// Nested loops, as in other programming languages
	// are loops within a loop. The first loop (outer)
	// prints a new line at the end of every sequence.
	// The second loop (inner) prints an asterisk for
	// every iteration.
	//

	total_lines := 20 // Total lines in sequence

	for counter := 0; counter < total_lines; counter++ { // Print new line every iteration

		for x := 0; x <= counter; x++ { // Print asterisks every iteration

			fmt.Print("*")

		}

		fmt.Println()

	}

}
