// Fibonacci closure https://tour.golang.org/moretypes/26
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {

	// initialise counters
	previous, latest := 0, 1

	// special case at the start
	start := 0

	// closure tracks the latest value printed and the previous value printed.
	// Each time the closure is called, it calculates the new value
	return func() int {
		// special case at the start
		switch start {
		case 0:
			// first run: move start index
			start = 1
			return 0
		case 1:
			// second run: no longer interested in start
			start = 2
			return 1
		}

		// new value
		new := latest + previous

		// advance the previous to the latest
		previous = latest

		// advance the latest to the new
		latest = new

		// return the new value
		return new
	}
}

func main() {
	f := fibonacci()
	for _ = range [10]int{} {
		fmt.Println(f())
	}
}

