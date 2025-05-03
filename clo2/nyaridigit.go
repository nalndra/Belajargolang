package main

import "fmt"

func main() {
	var x, n int
	fmt.Scan(&x, &n)

	cek := false

	for n > 0 {
		digit := n % 10
		if digit == x {
			cek = true
			break
		}
		n /= 10
	}

	fmt.Println(cek)
}
