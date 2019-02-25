

package main

import (
	"fmt"
)

// calculate the mean of two numbers 
func main() {
	var x float64  //2 types of floats, 32 and 64
	var y float64

	x = 1
	y = 2

	fmt.Printf("x=%v, type of %T\n", x, x)
	//fmt.Printf function gets a template to print and a value to fill the function
	//%v verb prints a go object
	//%t verb prints it's type
	fmt.Printf("y=%v, type of %T\n", y, y)

		//define mean
	var mean float64
		//logic to calculate mean
	mean = (x + y) / 2.0
		//template to print and it's value/types
	
	fmt.Printf("result: %v, type of %T\n", mean, mean)
}
