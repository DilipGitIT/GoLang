package main

import (
	"fmt"
)

func main() {
	var num1, num2, result1 int
	num3 := 20
	fmt.Printf("Enter the values")
	fmt.Scanf("%d %d", &num1, &num2)
	result1 = num1 + num2 + num3
	fmt.Printf("Result= %v", result1)

}
