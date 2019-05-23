package main

import (
	"fmt"
	"math/big"
)

// We can rewrite our prime/even/odd number
// by using a switch statement
//
//

func main() {

	counter := 1
	iterationLimit := 100

	for counter <= iterationLimit {

		isPrime := big.NewInt(int64(counter)).ProbablyPrime(2)
		isEven := true

		if counter%2 != 0 {
			isEven = false
		}

		switch { // Same as switch true

		case counter == 2:
			fmt.Println("Even Prime:", counter)
			counter++
			continue

		case !isPrime && isEven: // counter is prime AND even
			fmt.Println("Even:", counter)
			counter++
			continue

		case isPrime && !isEven: // counter is prime NOT even
			fmt.Println("Odd Prime:", counter)
			counter++
			continue

		default: // counter not prime is odd
			fmt.Println("Odd:", counter)
			counter++
			continue

		}

	}

}
