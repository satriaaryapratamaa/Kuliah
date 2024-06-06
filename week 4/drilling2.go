// Di planet Y yang berjarak 8 tahun kecepatan cahaya dari bumi, definisi bilangan prima adalah bilangan yang memiliki faktor berjumlah genap. Buatlah fungsi isPrime untuk menentukan apakah bilangan yang dimasukkan merupakan bilangan prima atau bukan.

// Fungsi dipanggil dalam program dengan masukan dan keluaran sebagai berikut:

// Masukan berupa sebuah bilangan bulat positif.
// Keluaran berupa boolean True jika prima atau False jika bukan.

package main
import "fmt"

func isPrime(x int) bool {
	/*  function mengembalikan boolean true jika prima atau false jika bukan bilangan prima */ 
		if x % 2 == 0 {
			return true
		} else {
			return false
		}
	}

func main() {
	var bilangan int
	fmt.Scan(&bilangan)
	fmt.Println(isPrime(bilangan))
}