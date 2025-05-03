package main
import "fmt"

type menabung struct {
    nama string
    persentase, masuk, saldo float64
}

func main() {
    var tabungan menabung

    fmt.Scan(&tabungan.nama, &tabungan.persentase, &tabungan.masuk, &tabungan.saldo)
    if tabungan.masuk*tabungan.persentase/100 == tabungan.masuk-tabungan.saldo {
        fmt.Printf("%s, kamu berhasil menabung sesuai target, 30%% dari pendapatan kamu\n", tabungan.nama)
    } else if tabungan.masuk*tabungan.persentase/100 > tabungan.masuk-tabungan.saldo {
        fmt.Printf("%s, kamu berhasil menabung melebihi target, 30%% dari pendapatan kamu\n", tabungan.nama)
    } else {
        fmt.Printf("%s, kamu gagal menabung sesuai target, 30%% dari pendapatan kamu\n", tabungan.nama)
    }
}
