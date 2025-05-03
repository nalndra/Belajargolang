package main

import "fmt"

func main() {
	var n, duaN, finish, scoreOne, scoreTwo int
	var timOne, timTwo, namaTim, winner string
	fmt.Scan(&n, &timOne, &timTwo)

	duaN = 2 * n

	for finish < duaN {
		fmt.Scan(&namaTim)
		finish++

		if namaTim == timOne && finish == 1 {
			scoreOne += 7
		} else if namaTim == timOne && finish == 2 {
			scoreOne += 6
		} else if namaTim == timOne && finish == 3 {
			scoreOne += 5
		} else if namaTim == timOne && finish == 4 {
			scoreOne += 3
		} else if namaTim == timOne && finish == 5 {
			scoreOne += 1
		} else if namaTim == timTwo && finish == 1 {
			scoreTwo += 7
		} else if namaTim == timTwo && finish == 2 {
			scoreTwo += 6
		} else if namaTim == timTwo && finish == 3 {
			scoreTwo += 5
		} else if namaTim == timTwo && finish == 4 {
			scoreTwo += 3
		} else {
			scoreTwo += 1
		}
	}

	if scoreOne > scoreTwo {
		winner = timOne
	} else {
		winner = timTwo
	}
	fmt.Println("tim", timOne, "memperoleh poin", scoreOne)
	fmt.Println("tim", timTwo, "memperoleh poin", scoreTwo)
	fmt.Println("pemenang adalah tim", winner)
}
