package my_function

import (
	"fmt"

	"github.com/tawan1314/go_testing/my_function/internal/sum"
)

func MyFunction(number1, number2 int) {
	fmt.Println("Hello My Function")
	result := sum.SumTwoNums(number1, number2)
	fmt.Println("Result is ", result)
}
