// Buatlah prosedur rekursif bernama reverse untuk mencetak string secara terbalik.

// Masukan berupa sebuah string

// Keluaran mencetak string secara terbalik

// Note: n dalam parameter prosedur adalah panjang dari string

package main

import "fmt"

func reverse(str string, n int) {
	/*{ I.S. Terdefinisi str sebagai input string
	F.S. menampilkan string yang terbalik menggunakan fungsi rekursif*/
	if n > 0 {
		var temporary string
		temporary = reverse(n-1)
		fmt.Print(temporary)
	}
}

func main() {
	var str string
	var n int
	fmt.Scan(&str)
	n = len(str)
	reverse(str, n)
}
