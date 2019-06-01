package Arrays

import "fmt"

// Slicing in Go works much like it does in Python
// You can select portions of an array by using the
// notation
//
// array := [5]int{1,2,3,4,5}
// var slice []int = array[2:3] <- Creates slice from array

func main() {
	// Create array
	array1 := [5]int{1, 2, 3, 4, 5}
	slice := array1[2:3]

	fmt.Println(slice)

}
