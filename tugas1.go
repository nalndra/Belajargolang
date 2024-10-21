package main

import (
	"fmt"
)

func f(x, y float64) float64 {
	return (5 / (2*y + 7)) + x*x - 3*y
}

func main() {
	var x, y float64

	fmt.Print("Nilai (x y) yang diinginkan: ")
	fmt.Scan(&x, &y)

	hasil1 := f(x, y)
	hasil2 := f(y, x)

	fmt.Printf("hasil untuk f(x, y) = %.4f\n", hasil1)
	fmt.Printf("hasil untuk f(y, x) = %.4f\n", hasil2)
}
