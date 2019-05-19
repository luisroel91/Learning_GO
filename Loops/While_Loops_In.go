package main

import "fmt"
import "math/big"

// There is only one kind of loop in Go.
// The "For" loop...but there is a workaround
//
// We can write the loops in alternate formats
// to accomplish loop constructs available in
// other languages
//

func main() {

	counter := 0           // Init counter
	iterationLimit := 1000 // Set this to you own value before running

	for counter <= iterationLimit {
		// Here we're creating a big.Int type var which
		// takes our counter casted into an Int64
		//
		// We the use the ProbablyPrime method which returns false/true
		// depending on whether our counter is prime or not
		//

		isPrime := big.NewInt(int64(counter)).ProbablyPrime(2)
		isEven := true

		if counter%2 != 0 {
			// Set isEven to false if number cannot be evenly divided by two
			isEven = false
		}

		if counter == 2 { // If prime AND even (two is the only even prime number)

			fmt.Println("Even Prime:", counter)
			counter++
			continue

		} else if !isPrime && isEven { // If NOT prime AND even
			fmt.Println("Even:", counter)
			counter++
			continue

		} else if isPrime && !isEven { // If prime AND NOT even
			fmt.Println("Odd Prime:", counter)
			counter++
			continue

		} else if !isPrime && !isEven { // If NOT prime AND NOT even
			fmt.Println("Odd:", counter)
			counter++
			continue
		}

	}
}
