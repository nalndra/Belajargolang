package main
import "fmt"

const NMAX int = 1024
type infoBuku struct {
    idBuku, judulBuku, penulisBuku string
}

type arrBuku[NMAX] infoBuku

func pencarianBuku(dataBuku arrBuku, nBuku int, searchKey string){
    var i int
    var ketemu bool = false
	for i = 0; i < nBuku; i++ {
		if dataBuku[i].judulBuku == searchKey {
			ketemu = true
			break
		}
	}

    if ketemu {
		fmt.Println("ID:", dataBuku[i].idBuku)
		fmt.Println("Judul Buku:", dataBuku[i].judulBuku)
		fmt.Println("Penulis Buku:", dataBuku[i].penulisBuku)
    } else {
		fmt.Println("Buku Tidak Tersedia")
    }
}
func main(){
    var dataBuku arrBuku
    var nBuku, i int
    var searchKey string

    fmt.Scan(&nBuku)
    for i = 0; i < nBuku; i++ {
		fmt.Scan(&dataBuku[i].idBuku, &dataBuku[i].judulBuku, &dataBuku[i].penulisBuku)
    }
    fmt.Scan(&searchKey)
    pencarianBuku(dataBuku, nBuku, searchKey)
}