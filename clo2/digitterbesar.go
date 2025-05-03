package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	maxDigit := 0 

	
	for n > 0 {
		digit := n % 10 
		if digit > maxDigit {
			maxDigit = digit 
		}
		n /= 10
	}

	fmt.Println(maxDigit)
}