package main

import "fmt"

type Klub struct {
	Nama         string
	Point        int
	SelisihGol   int
}

func main() {
	var klubA, klubB, klubC Klub

	fmt.Scan(&klubA.Nama)
	fmt.Scan(&klubA.Point)
	fmt.Scan(&klubA.SelisihGol)

	fmt.Scan(&klubB.Nama)
	fmt.Scan(&klubB.Point)
	fmt.Scan(&klubB.SelisihGol)
	
	fmt.Scan(&klubC.Nama)
	fmt.Scan(&klubC.Point)
	fmt.Scan(&klubC.SelisihGol)

	var terbanyak Klub
	var selisih Klub

	if klubA.Point > klubB.Point && klubA.Point > klubC.Point {
		terbanyak = klubA
	} else if klubB.Point > klubA.Point && klubB.Point > klubC.Point {
		terbanyak = klubB
	} else {
		terbanyak = klubC
	}

	if selisih.SelisihGol < klubA.SelisihGol && (klubA.Point == terbanyak.Point) {
		selisih = klubA
	} else if selisih.SelisihGol < klubB.SelisihGol && (klubB.Point == terbanyak.Point) {
		selisih = klubB
	} else if selisih.SelisihGol < klubC.SelisihGol && (klubC.Point == terbanyak.Point) {
		selisih = klubC
	}

	fmt.Println(terbanyak.Nama, selisih.Nama)
}