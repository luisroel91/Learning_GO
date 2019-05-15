package main

import "fmt"

func main() {
	var age int // declare int variable

	fmt.Println("my age is", age)

	age = 28 // Assign value

	fmt.Println("my age is", age)

	age = 54

	fmt.Println("my new age is", age)

	var text string = "Hello World" //declare string var

	fmt.Println(text)

	var variable = 2.4 // Type inference dictates this will be a float64

	fmt.Printf("%T\n", variable)

}
