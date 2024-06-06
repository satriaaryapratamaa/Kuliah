// Buatlah fungsi dalam bahasa Go bernama kelilingPersegi untuk menghitung keliling persegi. Fungsi memiliki sebuah parameter berupa sisi persegi dengan tipe bilangan bulat dan mengembalikan nilai berupa bilangan bulat.

// Fungsi akan dipanggil dalam sebuah program dengan masukan dan keluaran sebagai berikut:

// Masukan berupa sebuah bilangan bulat yang menyatakan sisi persegi.
// Keluaran adalah keliling persegi dengan sisi sesuai masukan. 

package main
import "fmt"

func kelilingPersegi(x int) int {
    /* fungsi mengembalikan keliling persegi */
	var hasilKeliling int

	hasilKeliling = x * 4
	return hasilKeliling
}

func main(){
	var sisi, keliling int
	fmt.Scan(&sisi)
	keliling = kelilingPersegi(sisi)
	fmt.Println(keliling)
}