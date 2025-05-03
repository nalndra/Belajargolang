package main

import "fmt"

func main() {
	var x int
	var biner string 
	fmt.Scan(&x)

	for x > 0 {
		if x % 2 == 0 {
			biner = "0" + biner 
		} else {
			biner = "1" + biner 
		}
		x /= 2 
	}

	fmt.Println(biner)
}