package main

import (
	"fmt"
)

func main() {
	var operand1, operand2, hasil float64
	var operator rune

	fmt.Scanf("%f %c %f", &operand1, &operator, &operand2)


	switch operator {
	case '+':
		hasil = operand1 + operand2
		fmt.Println(hasil)
	case '-':
		hasil = operand1 - operand2
		fmt.Println(hasil)
	case '*':
		hasil = operand1 * operand2
		fmt.Println( hasil)
	case '/':
			hasil = operand1 / operand2
			fmt.Println("%.3f",hasil)
	}
}
