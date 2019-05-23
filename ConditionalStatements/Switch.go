package main

import "fmt"

// Helper function for below
// Takes a number, multiplies it by 15
func number(factor int) int {
	mulNum := 15 * factor
	return mulNum
}

// Switches are a more concise way or writing
// if else statements. When a switch runs, it will
// check each case to see if there is a match.
// If there is, it will execute the code

func main() {
	letter := "i"

	switch letter {
	case "a", "e", "i", "o", "u": //multiple expressions in case
		fmt.Println("vowel")
	default:
		fmt.Println("not a vowel")
	}

	switch finger := 4; finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default: //default case
		fmt.Println("incorrect finger number")
	}

	switch num := 50; {
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200\n", num)
	}

	// Note that we can generate our expression at runtime
	// we do it below by calling the number function we
	// defined at the top of this file
	//
	// The fallthrough keyword helps us run code in the
	// same way elif helps us in Python
	//
	// Typically, as soon as a case is matched, the program
	// will break out of the switch. With fallthrough, we
	// keep running even after the first

	switch swNum := number(4); {
	case swNum < 50:
		fmt.Printf("swNUM:%d is lesser than 50\n", swNum)
		fallthrough
	case swNum < 100:
		fmt.Printf("swNUM:%d is lesser than 500\n", swNum)
		fallthrough
	case swNum < 200:
		fmt.Printf("swNUM:%d is lesser than 1000", swNum)
	}

}
