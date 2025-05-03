package main
import "fmt"

const NMAX = 10

type tabInt [NMAX]int

var data tabInt
var nData int

func baca() {
	var num int
	for {
		fmt.Scan(&num)
		if num == 0 || nData >= NMAX {
			break
		}
		if num < 0 {
			num = -num
		}
		data[nData] = num
		nData++
	}
}

func cetak() {
	for i := 0; i < nData; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(data[i])
	}
	fmt.Println()
}

func jumlah() int {
	sum := 0
	for i := 0; i < nData; i++ {
		sum += data[i]
	}
	return sum
}

func rataRata() float64 {
    if nData == 0 {
        return 0
    }
    return float64(jumlah()) / float64(nData)
}

func main() {
	baca()
	cetak()
	
	fmt.Println(jumlah())
	fmt.Printf("%.1f\n", rataRata())
}
