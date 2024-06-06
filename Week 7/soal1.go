package main

import "fmt"

const NMax int = 1000

type tabstr [NMax]string
type tabint [NMax]int

func main() {
	var n, i, totalP int
	var rata2 float64
	var nm tabstr
	var FT, TW, TH tabint

	//masukan banyaknya pemain
	fmt.Scan(&n)

	//masukan nama, banyaknya poin yang didapat. menghitung total poin per pemain n.
	for i = 0; i < n; i++ {
		fmt.Scan(&nm[i], &FT[i], &TW[i], &TH[i])
	}

	//Menampilkan poin yang diperoleh per-pemain dan menjumlah seluruh total poin pemain
	for i = 0; i < n; i++ {
		total := FT[i] + (TW[i] * 2) + (TH[i] * 3)
		totalP += total
		fmt.Println("Poin yang diperoleh", nm[i], "adalah", total)
	}

	//Menghitung dan menampilkan rata-rata poin yang didapat seluruh pemain.
	rata2 = float64(totalP) / float64(n)
	fmt.Printf("Rata-rata poin yang diperoleh seluruh pemain adalah %.2f", rata2)
}
