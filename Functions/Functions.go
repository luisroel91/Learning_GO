package Functions

// Lets declare a function
// Function format is as follows:
//
// func functionname(parametername type) returntype {
//   function body
// }

// PS: The /n in front of our format strings
// is to make the console output on a new line

import "fmt"

func add(x int, y int, z int) int {
	return x + y + z
}

// We can avoid having to declare types for our parameters
// if all parameters are of the same type

func add2(x, y, z int) int {
	return x + y + z
}

// We can also return multiple values

func CalcArea(length, width int) (int, int, int) {
	area := length * width

	return length, width, area
}

// Or, if we can use named return values

func CalcArea2(length, width int) (_len, _width, area int) {
	_len = length
	_width = width

	area = length * width

	return // No explicit return
}

func Functions() {
	variable := add(10, 5, 20)

	variable2 := add2(20, 12, 1)

	length, width, area := CalcArea(10, 50)

	fmt.Println(
		"Return:", variable,
		"Return2:", variable2,
	)

	fmt.Printf(
		"Len:%v\nWidth:%v\nArea:%v", length, width, area,
	)

	// We can also use the blank identifier if we don't need
	// everything our function returns

	area1, _, _ := CalcArea2(10, 23)

	fmt.Println(
		"\nArea1:",area1,
	)

}
