package main
import "fmt"
func main(){
	var nopol, i int
	fmt.Scan(&nopol)

	for nopol >= 10 {
		i = 0
		for nopol > 0 {
			i += nopol % 10
			nopol /= 10
		}
		nopol = i
	}
	fmt.Print (i)
}