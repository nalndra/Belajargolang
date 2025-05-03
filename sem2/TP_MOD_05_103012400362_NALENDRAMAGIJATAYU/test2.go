package main
import "fmt"

func mtk(a, b int) int {
    if b == 0 {
		return a
    }else{
		return mtk(b, a % b)
	}
}

func main() {
    var x, y, hasil int
	fmt.Scan(&x, &y)

    hasil = mtk(x, y)

    fmt.Println(hasil)
}
