package main

import "fmt"

func main() {
	// for loop
	// for x := 0; x < 10; x++ {
	// 	fmt.Println(x)
	// }

	// example
	var productList = []string{"Macbook", "iPhone", "iPod", "AirPod"}
	for index, product := range productList {
		fmt.Println(index, product)
	}

	// do while loop
	// y := 0
	// for {
	// 	fmt.Println(y)
	// 	y++
	// 	if y >= 10 {
	// 		break
	// 	}
	// }

	// example
	for index, i := range productList {
		fmt.Println(index, i)
		if i == "iPod" {
			break
		}
	}

	// while loop
	// z := 0
	// for z < 10 {
	// 	fmt.Println(z)
	// 	z++
	// }
}
