package main

import "fmt"

func main() {
	customerInfo := make(map[string]string)
	customerInfo["name"] = "Tawan"
	customerInfo["age"] = "24"
	customerInfo["sex"] = "Male"
	customerInfo["career"] = "Software Tester"
	fmt.Println(customerInfo)

	for key, info := range customerInfo {
		fmt.Printf("%s -> %s\n", key, info)
	}
}
