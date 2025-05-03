package main
import "fmt"

const NMAX int = 1024
type arrBilangan[NMAX] int

func main(){
    var bilBul arrBilangan
    var nBilangan, bilangan int

    fmt.Scan(&bilangan)
    for bilangan != 0 {
        bilBul[nBilangan]= bilangan
        nBilangan = nBilangan + 1
        fmt.Scan(&bilangan)
    }
    fmt.Println(cariNilaIMax(bilBul, nBilangan), cariNilaIMin(bilBul, nBilangan))
}

func cariNilaIMax(arrayBilbul arrBilangan, nBilangan int)int{
    var i int
    var maxIndex int = arrayBilbul[0]

    if nBilangan == 0 {
        return -1
    }
    for i = 1; i < nBilangan; i++ {
        if arrayBilbul[i] > maxIndex {
            maxIndex = arrayBilbul[i]
        }
    }
    return maxIndex
}

func cariNilaIMin(arrayBilbul arrBilangan, nBilangan int)int{
    var i int
    var minIndex int = arrayBilbul[i]

    if nBilangan == 0 {
        return -1
    }
    for i = 1; i < nBilangan; i++ {
        if arrayBilbul[i] < minIndex {
            minIndex = arrayBilbul[i]
        }
    }
    return minIndex
}