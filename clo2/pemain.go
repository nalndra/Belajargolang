package main

import "fmt"

func main() {
	var lv int
	var stat string
	fmt.Scan(&lv)

	switch {
	case lv >= 1 && lv <= 10:
		stat = "pemula"
	case lv >= 11 && lv <= 20:
		stat = "menengah"
	case lv >= 21 && lv <= 30:
		stat = "ahli"
	case lv > 30:
		stat = "master"
	default:
		stat = "Level tidak valid"
	}
	fmt.Println(stat)
}	