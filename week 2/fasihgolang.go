// Di planet X yang berjarak 30 tahun cahaya dari planet Bumi syarat untuk mendapatkan KTP adalah minimal berumur 10 tahun dan fasih berbahasa golang. Buatlah fungsi dalam bahasa GO bernama dapatKTP untuk menentukan apakah penduduk planet X pantas mendapat KTP atau tidak.

// Masukan berupa sebuah bilangan bulat yang menyatakan umur dan sebuah boolean yang menyatakan status kefasihan berbahasa golang (True bila fasih atau False bila tidak).
// Keluaran berupa boolean True jika bisa mendapatkan KTP, atau False jika belum bisa.

package main 
import "fmt"

func dapatKTP(umur int, golang bool) bool {
	// mengembalikan boolean true jika umur lebih besar dari atau sama dengan 10
	// dan fasih berbahasa golang
		
		golang = true 
		golang = umur >= 10 
		return golang
	
	}

func main() {
	var umur int
	var golang bool
	fmt.Scan(&umur, &golang)
	fmt.Println(dapatKTP(umur, golang))
}