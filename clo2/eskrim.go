package main
import "fmt"
func main () {
	var es, x, susu int 
	
	fmt.Scan(&es, &x)
	susu = (es * x) / 1000
	if (es * x) % 1000 != 0 {
		susu = susu + 1
	}
	fmt.Printf("%d karung", susu)
}
