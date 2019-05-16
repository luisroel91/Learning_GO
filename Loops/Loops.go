package main

import "fmt"

func main() {

	// Here we'll define our loop. The syntax for a standard loop is:
	//
	// for initial var; exit condition; operation per iter {
	//
	// }
	//
	// In the below loop, counter starts at 1 and every iteration adds 1
	// to the counter. The loops runs until the counter is greater than
	// or equal to 500. We use an if statement to break the loop when the
	// counter reaches 201. If its an even number, it notifies the user
	// and continues the loop using the continue keyword.

	for counter := 1; counter <= 500; counter++ {
		fmt.Println(
			"Counter is:", counter,
		)

		if counter >= 201 {
			break // loops terminates when counter reaches 200
		}

		if counter%2 == 0 {
			fmt.Println(
				"Counter is an even number:", counter,
			)
			continue //continue
		}

	}

}
