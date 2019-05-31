package main

import "fmt"

// Arrays belong to type n[T] where n is the number
// of elements in the array and T is the type of those
// elements.
//
// Arrays are VALUE types not REFERENCE types
// Meaning that when an array is assigned to a new variable
// a copy of the original array is assigned to the new variable
// not a reference to the old one

func main() {
	// var a [7]int creates an empty array with 4 zeros
	var a [4]int

	// We can then insert specific values into the array
	// using slicing, sort of like in Python. Also, like
	// Python, arrays start a zero

	a[0] = 1
	a[1] = 5
	a[2] = 10
	a[3] = 15

	// But this is tedious, more than likely, you're going to
	// need to create an array with values or variables already
	// contained in said array. We can do that shorthand using
	// the following format:
	//
	// variable name := [size of array]type of array{value, value, value}
	//
	// KEEP IN MIND, that the size is PART of the type. Meaning:
	//
	// arr0 := [3]int{1,2,3}
	// var arr1 [5]int
	// b = a <- Not possible, because technically, they are diff types
	//
	// Error "Cannot use a (type [3]int) as type [5]int in assignment

	b := [4]int{20, 25, 30, 35}

	// All values are not required when assigning
	// For example, if an array is size 4, we don't
	// always have to input all 4 values when creating
	// the array

	c := [4]string{"Hello", "World"}

	// We can also ignore the size parameter instead
	// using ... and the array length will automatically
	// be determined by the compiler

	d := [...]int{40, 45, 50, 55, 60, 65, 70}

	// Again, arrays are value types, not references

	e := d
	e[0] = 100

	// Output
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	// Output length, notice how the array is automatically sized
	fmt.Println(len(d))

	// Original array has not been changed
	fmt.Println(d)
	// Array E has been changed
	fmt.Println(e)

	// One way to loop over an array's values and indexes
	// are by using its length in a standard for loop.
	for i := 0; i < len(a); i++ {
		fmt.Printf("Element %d of Arr A is %d\n", i, a[i])
	}

	// Another, shorter more concise way is by using the range
	// form of the for loop. This is similar to using enumerate()
	// in Python. Note you do not have to specifically use
	// the words "index", "value". You can use any variable name
	// you'd like

	accumulator := int(0)

	for index, value := range b {
		fmt.Printf("Element %d of Arr B is %d\n", index, value)
		// Add every value to the accumulator
		accumulator += value
	}

	// Print sum of Array B
	fmt.Println("Sum of all elements is ", accumulator)

	// You can also ignore the index by using the blank
	// identifier "_"
	for _, value := range d {
		fmt.Println("Value of d is ", value)
	}
}
