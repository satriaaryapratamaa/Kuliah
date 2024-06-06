package main

import "fmt"

func segitigaRecursive(baris int, kolom int, n int) {
	/*  I.S Terdefinisi nilai bilangan bulat baris, kolom, dan n
	F.S menampilkan pola  string "*" yang berbentuk segitiga */
	if kolom <= n {
		if kolom <= n - baris { // Pengecekan Pengulangan kolom dan baris dengan operasi perbandingan,
			fmt.Print("")
		} else {
			fmt.Print("*")
		}
		segitigaRecursive(baris, kolom+1, n)
	} else if baris < n {
		fmt.Println()
		segitigaRecursive(baris+1, 1, n)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	segitigaRecursive(1, 1, n)
}
