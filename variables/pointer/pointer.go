package main

import (
	"fmt"
)

func zerovalue(x int) {
	x = 0
}

func zeropointer(x2 *int) {
	*x2 = 0
}

func main() {
	i := 1
	fmt.Println("i = ", i)

	zerovalue(i)
	fmt.Println("i value from zerovalue = ", i)

	zeropointer(&i)
	fmt.Println("i pointer from zeropointer = ", &i)
}
