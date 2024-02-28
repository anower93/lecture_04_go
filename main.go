package main

import (
	"fmt"
)

func main() {

	//Loop, In Go programming it has only one loop

	//Simple for loop
	for i := 1; i <= 9; i++ {
		fmt.Println(i)

	}

	// Range in for loop

	students := [...]string{"One", "Two", "Three"}

	for i, std := range students {
		fmt.Println(i, std)
	}

	//For while loop

	x := 0
	for x < 11 {
		fmt.Println(x, "Hello")
		x++
	}

	// For infinity

	for {
		fmt.Println("Hello")
		x++
	}

}
