package main
import "fmt"

const NMAX int = 1024
type arrinteger [NMAX] int

func cariNila1Min(A arrinteger, size int)int{
    var i int
    var minIndex int = A[i]
	
    if size == 0 {
		return -1
    }
    for i = 1; i < size; i++{
		if A[i] < minIndex {
			minIndex = A[i]
		}
    }
    return minIndex
}

func main() {
    var A arrinteger
    var size, i int
	
    fmt.Scan(&size)
    for i = 0; i < size; i++{
		fmt.Scan(&A[i])
    }
    fmt.Println(cariNila1Min(A,size))
}