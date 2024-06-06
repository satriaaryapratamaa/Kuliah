// Buatlah function dalam bahasa Go untuk menjumlahkan dua bilangan ganjil. 
//Fungsi memiliki dua parameter berupa bilangan bulat yang akan dijumlahkan dan 
//mengembalikan nilai berupa bilangan bulat hasil penjumlahan jika kedua bilangan itu ganjil.
//Jika tidak ganjil keduanya akan dikembalikan nilai 0. Selanjutnya fungsi akan dipanggil dalam sebuah program dengan masukan dan keluaran sebagai berikut:
// Masukan terdiri dari dua bilangan bulat yang akan dijumlahkan.
// Keluaran adalah hasil penjumlahan dua bilangan bulat itu, jika keduanya ganji. Jika tidak ganjil keduanya, maka hasil penjumlahannya bilanga 0.

package main

import "fmt"

func hitungJumlah(x, y int) int {
	/* function mengembalikan jumlah dua bilangan, jika kedua bilangan itu ganjil
	Jika tidak ganjil keduanya, kembalikan bilangan 0
	*/
	//var hasil int

	if x % 2 == 0 {
		return 0
	} else if y % 2 == 0 {
		return 0
	} else {
		return x + y
	}
	
}

func main() {
	var b1, b2 int
	fmt.Scan(&b1, &b2)
	fmt.Println(hitungJumlah(b1, b2))
}
