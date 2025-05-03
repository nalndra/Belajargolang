package main

import "fmt"

func main() {
	var n, totalmobil, finish int
	var x, y, mobilfinish, pemenang string
	var skorx, skory int

	fmt.Scan(&n, &x, &y)

	finish= 0
	totalmobil= 2*n
	skorx = 0
	skory = 0

	for finish < totalmobil{
		fmt.Scan(&mobilfinish)
		finish++

		if mobilfinish == x && finish == 1{
			skorx += 7
		}else if mobilfinish == x && finish == 2 {
			skorx += 6
		}else if mobilfinish == x && finish == 3 {
			skorx += 5
		}else if mobilfinish == x && finish == 4 {
			skorx += 3
		}else if mobilfinish == x && finish == 5 {
			skorx += 1
		}else if mobilfinish == y && finish == 1 {
			skory += 7
		}else if mobilfinish == y && finish == 2 {
			skory += 6
		}else if mobilfinish == y && finish == 3 {
			skory += 5
		}else if mobilfinish == y && finish == 4 {
			skory += 3
		}else{
			skory += 1
		}
	}
	if skorx > skory {
		pemenang = x
	}else {
		pemenang = y
	}
	fmt.Println("tim", x, "memperoleh poin", skorx)
	fmt.Println("tim", y, "memperoleh poin", skory)
	fmt.Println("pemenang adalah tim", pemenang)
}