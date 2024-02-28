package main

import "fmt"

func main() {

	fmt.Println("Enter you age: ")

	var age int

	fmt.Scanf("%d", &age)

	if age < 3 && age > 0 {
		fmt.Println("Infant")
	} else if age >= 3 && age < 13 {
		fmt.Println("Children")
	} else if age >= 13 && age < 18 {
		fmt.Println("Young man!")
	} else if age >= 18 && age < 60 {
		fmt.Println("Adult")
	} else if age >= 60 {
		fmt.Println("Old man!")
	} else if age < 0 {
		fmt.Println("Something Went Wrong!")
	}

	switch age {
	case 2:
		fmt.Println("Your are 2 years old!")
	case 3:
		fmt.Println("You are 3 years old!")
	case 4:
		fmt.Println("You are 4 years old!")
	case 5:
		fmt.Println("You are 5 years old")
	default:
		fmt.Println("Something went wrong!")
	}

}
