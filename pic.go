// Pic https://tour.golang.org/moretypes/18
package main

import "golang.org/x/tour/pic"

// picFunc returns the value of the grid, given its coordinates
type picFunc func(x, y int) int

// picMaker is passed to pic.Show
type picMaker func(dx, dy int) [][]uint8

// picWith converts a picFunc to a picMaker
// (that can be passed to pic.Show), ideally pic.Show would just take the picFunc
func picWith(fn picFunc) picMaker {

	// return a function that produces a matrix
	return func(dx, dy int) [][]uint8 {

		// initialise
		matrix := make([][]uint8, dy)

		for y := range matrix {
			// initialise
			row := make([]uint8, dx)

			for x := range row {
				// calc cell value, convert type
				row[x] = uint8(fn(x, y))
			}

			// set completed row to the matrix
			matrix[y] = row
		}

		return matrix
	}
}

func showPic(fn picFunc) {
	// composition of Show and picWith
	pic.Show(picWith(fn))
}

// one return 1, just so we can test
func one(x, y int) int {
	return 1
}

// times: x * y
func times(x, y int) int {
	return x * y
}

// pow: x ^ y
func pow(x, y int) int {
	return x ^ y
}

// square: x ^ 2
func square(x, y int) int {
	return x ^ 2
}

func main() {
	// anonymous func
	showPic(func(x, y int) int {
		// plus one to avoid divide by zero
		return (x * y) / ((x ^ y) + 1)
	})

	// show pics for various picFuncs
		showPic(pow)
		showPic(square)
		showPic(times)
}

