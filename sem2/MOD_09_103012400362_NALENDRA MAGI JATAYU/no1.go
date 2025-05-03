package main
import "fmt"

const NMAX int = 1024
type arrPenjualan [NMAX] int

func nilaiMax(A arrPenjualan, size int)int{
	var maxIndex int = 0
	var i int
	
	if size == 0 {
		return -1
	}
	
	for i = 1; i < size; i++{
		if A[i] > A[maxIndex]{
			maxIndex = i
		}
	}
	return maxIndex
}

func main(){
	var A arrPenjualan
	var bilbul, i, size, idxMax int
	
	i = 1
	size = 1
	fmt.Scan(&bilbul)
	for bilbul != -1{
		A[i] = bilbul
		i = i + 1
		size = size + 1
		fmt.Scan(&bilbul)
	}
	idxMax = nilaiMax(A, size)
	fmt.Print(idxMax, A[idxMax])
	
}