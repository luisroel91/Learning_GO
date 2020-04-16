package main

// Now we'll take our util and use it to generate our PW's
// Remember, an ascii value is simply a number
import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
)

// We'll need two functions, one to generate
// our string of random bytes and one to encode
// it into a string.

// Takes in an int64 number and returns a slice of bytes
// if there is an error, we will return that as well
func generateRandomBytes(size int64) ([]byte, error) {
	// We use make() to create a new slice of bytes of size "size"
	randbytes := make([]byte, size)
	// Then we fill our "randbytes" with random numbers
	_, err := rand.Read(randbytes)
	// If there was an error reading into the slice
	if err != nil {
		// Notify user, return error
		fmt.Println("ERROR: ", err)
		return nil, err
	}
	// If there was no error, return our slice of random bytes
	return randbytes, nil
}

// Now we need a function that can take those bytes and
// encode them. It takes in a size, generates our random
// data and then returns a base64 encoded version of it
func generateString(size int64) (string, error) {
	randstr, err := generateRandomBytes(size)
	return base64.URLEncoding.EncodeToString(randstr), err
}

func main() {
	// We'll set this var to our desired PWLEN
	// whether it was input by the user or our defaulf of 8
	var PWLEN int64 = 0
	arguments := os.Args
	arglen := len(arguments)
	switch {
	// Arglen would ==2 with one arg
	// >=3 with two or more args
	// 1 if no args
	case arglen == 2:
		convInput, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			fmt.Println("ERROR: Input cannot be negative")
			os.Exit(2)
		}
		fmt.Printf("Using PW size of %d...", convInput)
		PWLEN = convInput
	case arglen >= 3:
		fmt.Println("ERROR: Too many arguments")
		os.Exit(2)
	default:
		fmt.Println("Using default PW size of 8...")
		PWLEN = 8
	}
	// Generate PW, our functionr returns an error
	// We need to handle it as well
	resultStr, err := generateString(PWLEN)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\nYour generated password is:")
	fmt.Println(resultStr[0:PWLEN])
}
