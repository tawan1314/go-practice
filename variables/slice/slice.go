package main

import "fmt"

func main() {
	var productPrice = []int{3000, 20000, 15000, 40000}
	fmt.Println(productPrice)

	// add values into the variable
	productPrice = append(productPrice, 5355)
	fmt.Println(productPrice)
}
