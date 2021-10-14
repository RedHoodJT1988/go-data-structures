package main

import "fmt"

// gets the power series of integer 'a' and returns tuple of square of 'a'
// and cube of 'a'
func powerSeries(a int) (square int, cube int) {
	square = a*a
	cube = square*a
	
	return square, cube
}

// main method
func main() {
	var square int
	var cube int
	square, cube = powerSeries(3)

	fmt.Println("Square ", square, "Cube", cube)
} 
