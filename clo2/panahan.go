package main

import "fmt"
func main(){
	
	var a, score, shots int

	for score < 30 {
		fmt.Scan(&a)
		score += a
		shots += 1
	}
fmt.Println(score, shots)
}