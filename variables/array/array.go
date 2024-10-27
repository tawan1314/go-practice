package main

import "fmt"

func main() {
	var productList = [4]string{"Macbook", "iPhone", "iPod", "AirPod"}
	fmt.Println(productList)

	// change a value into the variable
	productList[0] = "Android"
	fmt.Println(productList)

	// for loop in array
	for i := 0; i < len(productList); i++ {
		fmt.Println(productList[i])
	}
}
