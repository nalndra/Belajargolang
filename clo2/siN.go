package main

import "fmt"

func main() {
	var x, n int

	fmt.Print("Masukkan x dan N: ")
	fmt.Scanf("%d %d", &x, &n)

	if x%n == 0 {
		fmt.Printf("%d kelipatan %d? true\n", x, n)
	} else {
		fmt.Printf("%d kelipatan %d? false\n", x, n)
	}
}
