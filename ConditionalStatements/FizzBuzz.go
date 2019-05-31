package main

import "fmt"

func main() {

	counter := 1
	iterLimit := 100

	for counter <= iterLimit {

		switch {

		case counter%3 == 0 && counter%5 == 0:
			fmt.Println("FizzBuzz:", counter)
			counter++
			continue

		case counter%3 == 0:
			fmt.Println("Fizz:", counter)
			counter++
			continue

		case counter%5 == 0:
			fmt.Println("Buzz:", counter)
			counter++
			continue

		default:
			counter++
			continue

		}
	}

}
