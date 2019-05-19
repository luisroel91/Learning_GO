package main

import "fmt"

func main() {
	// We'll need this to check against
	// Shorthand declaration
	//

	number0 := 1
	number1 := 61

	// If statement, as in other languages
	// if condition is true, then code in
	// body runs
	//

	if number0 > 15 {
		fmt.Print("Number0 is greater than 15")
		// Similar to elif in Python
	} else if number1 < 50 {
		fmt.Print("Number0 is less than 50")
		// Below we use the && operator to combine two
		// checks into our conditional
	} else if number0 < 20 && number1 > 60 {
		fmt.Print("Both numbers passed the test")
	} else { // Else must be on same line as closing bracket due to how Go invisibly
		// inserts semicolons

	}

	number2 := 99

	// We can also declare the function in a way
	// that before the check is made, some code runs
	// if statement; condition {
	//
	// }
	//

	if number2++; number2 > 100 {
		fmt.Println("\nWe added 1 to the number2 var")
	}

}
