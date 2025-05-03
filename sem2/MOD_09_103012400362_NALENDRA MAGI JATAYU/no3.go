package main
import "fmt"

const NMAX int = 1024
type arrBaju [NMAX] string

func main () {
    var A arrBaju
    var warnaBaju, searchKey string
    var i, size int
	
    i = 0
    size = 1
    fmt.Scan(&warnaBaju)
    for warnaBaju != "." {
        A[i] = warnaBaju
        i = i + 1
        size = size + 1
        fmt.Scan(&warnaBaju)
    }
    fmt.Scan(&searchKey)
    cekBaju(A, size, searchKey)
}

func cekBaju(A arrBaju, size int, searchKey string){
    var i int
    var ketemu bool = false
    for i = 0; i < size; i++ {
        if A[i] == searchKey {
            ketemu = true
            break
        }
    }
    if ketemu {
        fmt.Println("Baju Ada")
    } else {
        fmt.Println("Baju Tidak Ada")
    }
}