// Buatlah prosedur rekursif lowercase menampilkan suatu string dengan semua hurufnya berupa huruf kecil

// Masukan suatu string dengan semua hurufnya kapital

// Keluaran menampilkan string dengan semua hurufnya berupa huruf kecil

package main

import "fmt"

// func lowercase(s string, n int) {
//     /*{ I.S. Terdefinisi s sebagai input string kapital 
// 	F.S. menampilkan string yang semua  huruf kecil menggunakan fungsi rekursif*/
// 	if n > 0 {
// 		fmt.Print(string(s[n-1] + 32))
// 		return lowercase(s string, n int)
// 	}
// }

// func main() {
// 	var str string
// 	fmt.Scanf("%s", &str)
	
// 	lowerStr := ""
// }

func lowercase(s string, n int) {
    /*{ I.S. Terdefinisi s sebagai input string kapital 
	F.S. menampilkan string yang semua  huruf kecil menggunakan fungsi rekursif*/
	if n > 0 {
		// masukkan fungsi rekursif pada baris ini
		lowercase(s, n-1)
		fmt.Print(string(s[n-1] + 32))
	}
}

func main () {
	var s string
	fmt.Scan(&s)
	lowercase(s, len(s))
}
