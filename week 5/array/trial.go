package main
import "fmt"

type tipe_buku struct {
	judul, penulis, penerbit string
	halaman, tahun_terbit int
}

func main() {
	var buku tipe_buku
	
	buku.halaman = 115
	buku.tahun_terbit = 2018
	
	fmt.Scan(&buku.judul, &buku.penulis, &buku.penerbit)
	fmt.Println("Jumlah halamannya adalah :", buku.halaman, "Tahun terbit :", buku.tahun_terbit)
}