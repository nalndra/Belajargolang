package main
import "fmt"

func main(){
	var a, b int
	var terkecil, terbesar, total int
fmt.Scan(&a, &b)

if a < b {
	terbesar, terkecil = b, a
}else {
	terbesar, terkecil = a, b
}
for x := terkecil; x <= terbesar; x++ {
	fmt.Print(x, " ")
}
fmt.Println()
total = 0
for x:= terkecil; x <= terbesar; x++ {
	total = total + x
}
fmt.Println("total: ", total)
}