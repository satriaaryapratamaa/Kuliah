package main
import "fmt"

func hitungSKS(totalSKS int) int {
    //  Hint :  1. Gunakan variabel lokal  yang tepat
    //          2. Gunakan fmt.scan untuk input kode mata kuliah berupa string 'A'-'E' 

	var kodeMatkul string
	fmt.Scan(&kodeMatkul)  // Scan Kode mata Kuliah

	if kodeMatkul == "0" {
		return totalSKS
	}

	sks := 0 // lakukan inisialisasi 
	
	if kodeMatkul[0] == 'A' || kodeMatkul[0] == 'B' || kodeMatkul[0] == 'C' {
		sks = 2
	} else if kodeMatkul[0] == 'D' || kodeMatkul[0] == 'E' || kodeMatkul[0] == 'F' {
		sks = 3
	}
	return hitungSKS(totalSKS + sks) // isi perhitungan penjumlahan SKS menggunakan fungsi rekursif
}


func main() {
	var totalSKS, sks int
	sks = 0
	totalSKS = hitungSKS(sks)
	fmt.Print(totalSKS)
}