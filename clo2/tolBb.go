package main

import "fmt"

func main() {
	
	countA, countB, countC := 0, 0, 0

	
	var input string
	for {
	
		fmt.Scan(&input)

		if input == "A" {
			countA++
		} else if input == "B" {
			countB++
		} else if input == "C" {
			countC++
		} else {
	
			break
		}
	}

	fmt.Printf("tipe A = %d\n", countA)
	fmt.Printf("tipe B = %d\n", countB)
	fmt.Printf("tipe C = %d\n", countC)
}
