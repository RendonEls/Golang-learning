

package main

import (
	"fmt"
)

// calculate the mean of two numbers 
func main() {

	x, y := 1.0, 2.0

	fmt.Printf("x=%v, type of %T\n", x, x)
	//fmt.Printf function gets a template to print and a value to fill the function
	//%v verb prints a go object
	//%t verb prints it's type
	fmt.Printf("y=%v, type of %T\n", y, y)

		//logic to calculate mean
	mean := (x + y) / 2.0
		//template to print and it's value/types
	
	fmt.Printf("result: %v, type of %T\n", mean, mean)
}
