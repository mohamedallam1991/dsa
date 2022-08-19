package main

import "fmt"

func powerSeries(a int) (square int, cube int, err error) {
	square = a * a
	cube = square * a
	//if we need to check for error, and return the error
	return square, cube, nil
}

func main() {
	var square int
	var cube int
	// The tuple
	square, cube, _ = powerSeries(3)

	fmt.Println("Square", square, "Cube", cube)
}
