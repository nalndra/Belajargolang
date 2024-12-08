package main
import "fmt"

func main(){
	var x, y, dif int
	var consecutive bool
	fmt.Scan(&x)

	y = x % 10
	x = x / 10

	dif = y - x % 10

	for x > 0 && dif == 1 || dif == -1 {
		y = x % 10
		x = x / 10
		dif = y - x % 10
	}
	consecutive = x == 0
	fmt.Println(consecutive)
}