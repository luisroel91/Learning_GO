package CalculateAreaOfSquare

import "fmt"

func CalculateAreaOfSquare() {
	// Declaring multiple vars

	var length, width = 100, 200

	// Calculated at runtime

	var area = length * width

	fmt.Println(
		"width:", width,
		"\nlength:", length,
		"\narea:", area,
	)

	// Declaring multiple vars of different types

	var (
		integer = 1
		text    = "Hello World"
		float   = 2.56
	)

	fmt.Println(
		"Integer:", integer,
		"\nString:", text,
		"\nFloat:", float,
	)

	// Shorthand declaration, singular

	integer2 := 12

	// Shorthand declaration, multiple
	// Note how you can declare different types
	// this way. It can only be used when at least
	// one of the variables on the left side is newly
	// declared

	integer3, name := 21, "Luis"

	fmt.Println(
		"\nInteger2:", integer2,
		"\nInteger3:", integer3,
		"\nName:", name,
	)

}
