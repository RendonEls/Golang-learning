
package main

import (
	"fmt"
)

// print numbers 1 to 20
// if the numbert is divisible by 3 and 5, print fizz buzz
// if the number is divisible by 3, print fizz
// if the number is divisible by 5, print buzz
// otherwise print the number 


func main() {
    fmt.Println("----")
    for i := 1; i <= 20; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("fizz buzz", i)
		} else if i % 3 == 0 {
            fmt.Println("fizz", i)
		} else if i % 5 == 0 {
			fmt.Println("buzz", i)
		} else {
			fmt.Println(i)
		}
	}
}