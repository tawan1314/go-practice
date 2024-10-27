package main

import (
	"fmt"

	"github.com/tawan1314/go_testing/my_function"
)

func main() {
	my_function.MyFunction(1, 2)

	// create a string variable
	var name string = "Tawan Duangchit"
	fmt.Println("My name is ", name)

	// create an integer variable
	var age int = 24
	fmt.Println("Age ", age)

	// create a float variable
	var score float32 = 98.99
	fmt.Println("Score ", score)

	// create a boolean variable
	var isPass bool = true
	fmt.Println("Exam pass ", isPass)

	// create a variable short form
	coffee := "Americano"
	fmt.Println("My favorite coffee is ", coffee)

	// display values in the variable
	var firstname string = "Tawan"
	fmt.Printf("Value is = %v\n", firstname)

	// display data types in the variable
	var surname string = "Duangchit"
	fmt.Printf("Data Type is %T", surname)
}
