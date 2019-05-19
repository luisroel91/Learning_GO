package main

import "fmt"

// We can use labels with the break keyword
// to break a particular loop if nested instead of
// all loops.
//
// This program is sorta like a gear. The outer loop
// continues until both counter_0 and counter_1 are
// the same. Then the inner loop break the outer using
// the break keyword and the "outer" label
//

func main() {
outer: // Label
	for counter_0 := 0; counter_0 < 10; counter_0++ {

		for counter_1 := 1; counter_1 < 12; counter_1++ {

			fmt.Printf("counter_0 = %d, counter_1 = %d\n", counter_0, counter_1)

			if counter_0 == counter_1 {
				fmt.Print("Breaking outer loop...")
				break outer

			}
		}
	}

}
