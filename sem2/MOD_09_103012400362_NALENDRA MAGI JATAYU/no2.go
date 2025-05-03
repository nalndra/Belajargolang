package main
import "fmt"

const NMAX = 1024
type satuanWaktu struct {
	menit, detik int
}

type arrWaktu [NMAX]satuanWaktu

func convertWaktu(waktuMenit, waktuDetik int) int {
	return waktuMenit*60 + waktuDetik
}

func waktuTercepat(waktuPembalap arrWaktu, nPembalap, menitMulai, detikMulai int) int {
	var selisihTerdekat, selisih int
	waktuMulai := convertWaktu(menitMulai, detikMulai)
	
	selisihTerdekat = convertWaktu(waktuPembalap[0].menit, waktuPembalap[0].detik) - waktuMulai
	
	for i := 1; i < nPembalap; i++ {
		selisih = convertWaktu(waktuPembalap[i].menit, waktuPembalap[i].detik) - waktuMulai
		if selisih < selisihTerdekat {
			selisihTerdekat = selisih
		}
	}
	return selisihTerdekat
}

func main() {
	var waktuPembalap arrWaktu
	var nPembalap, i, waktuMinimal int
	var menitMulai, detikMulai int
	
	fmt.Scan(&nPembalap)
	fmt.Scan(&menitMulai, &detikMulai)
	
	for i = 0; i < nPembalap; i++ {
		fmt.Scan(&waktuPembalap[i].menit, &waktuPembalap[i].detik)
	}
	
	waktuMinimal = waktuTercepat(waktuPembalap, nPembalap, menitMulai, detikMulai)
	fmt.Print(waktuMinimal/60, waktuMinimal%60)
}