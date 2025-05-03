package main
import "fmt"

const NMAX = 5

type tabInt struct {
	info [NMAX]int
	n int
}

func bacaData(A *tabInt) {
	if A.n >= NMAX {
		return
	}
	var num int
	fmt.Scan(&num)
	A.info[A.n] = num
	A.n++
}

func cetakData(A tabInt) {
	for i := 0; i < A.n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(A.info[i])
	}
	fmt.Println()
}

func main() {
	var data tabInt
	for i := 0; i < 6; i++ {
		bacaData(&data)
	}
	cetakData(data)
}