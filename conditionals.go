package main

import (
	"fmt"
)
// // working with if statement
// func main() {
// 	x := 10  //assign 10 to x

// 	if x > 5 { //is x bigger than 5?
// 		fmt.Println("x is big")
// 	}
// 	//what about when a condition is no met?

// 	if x> 100 {
// 		fmt.Println("x is HUGE")
// 	} else {
// 		fmt.Println("x is not as big")
// 	}
// 	if x > 5 && x < 15 {
// 		fmt.Println("x is perfect")
// 	}
// 	//double pipe for logical OR
// 	if x < 20 || x > 30 {
// 		fmt.Println("x is not in range")
// 	}
// 	a := 11.0
// 	b := 20.0
// 	// if 11 is more than half of 20 print something cool
// 	if frac := a / b; frac > 0.5 {
// 		fmt.Println("a is more than half of b")
// 	}
// }

func main() {
	x := 2
	
	switch x {
	case 1: 
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Printf("many")
	}
	
	switch {
	case x > 100:
		fmt.Println("x is very big")
	case x > 10:
		fmt.Println("x is big")
	default:
		fmt.Println("x is small")
	}
}