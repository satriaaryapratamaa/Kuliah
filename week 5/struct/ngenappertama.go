// Buatlah sebuah program dalam bahasa Go yang bertugas menghitung total dari n bilangan genap positif pertama dan 
//juga menentukan bilangan genap terakhir dalam rangkaian tersebut. Program ini harus menyimpan kedua informasi ini dalam sebuah struct.
// Sebagai contoh, jika n = 5, maka Anda perlu menghitung total dari 5 bilangan genap pertama, yaitu 2, 4, 6, 8, dan 10. Dalam kasus ini, 
//totalnya adalah 30 dan bilangan genap terakhir adalah 10. Kedua nilai ini harus disimpan dan ditampilkan dari dalam struct.
// Masukan Sebuah bilangan bulat n yang mewakili jumlah bilangan genap yang akan dihitung.
// Keluaran Bilangan genap terakhir dalam rangkaian dan total dari n bilangan genap pertama.\

package main 
import "fmt"

type type HasilGenap struct {
	BilanganTerakhir int
	JumlahTotal      int
}

func jumlahBilanganGenap(n int) HasilGenap {
	total := 0
	terakhir := 0
	for i := 1; n > 0; i++ {
		if i%2 == 0 {
			total += i
			terakhir = i
			n--
		}
	}
	return HasilGenap{BilanganTerakhir: terakhir, JumlahTotal: total}
}

func main() {
	var n int
	fmt.Scan(&n)

	hasil := jumlahBilanganGenap(n)
	fmt.Print(hasil.BilanganTerakhir)
	fmt.Print(" ", hasil.JumlahTotal)
}

//PENJELASAN!!!

// Di awal kode, kita mendeklarasikan sebuah struct bernama HasilGenap yang memiliki dua field: 
// BilanganTerakhir untuk menyimpan bilangan genap terakhir dalam rangkaian, dan JumlahTotal untuk menyimpan total dari n bilangan genap pertama.

// Fungsi jumlahBilanganGenap menerima parameter n yang merupakan jumlah bilangan genap yang akan dihitung.
// Kita inisialisasi variabel total dan terakhir dengan nilai awal 0. total akan digunakan untuk menghitung jumlah total dari n bilangan genap, sedangkan terakhir akan menyimpan bilangan genap terakhir dalam rangkaian.
// Di dalam loop for, kita iterasi dari 1 ke atas dan menghentikan iterasi ketika kita telah menemukan n bilangan genap.
// Setiap kali kita menemukan bilangan genap (dengan menggunakan operator modulus %), kita menambahkannya ke total, menyimpannya sebagai terakhir, dan mengurangi nilai n.
// Setelah loop selesai, kita kembalikan hasilnya dalam bentuk HasilGenap.

// Fungsi main adalah entry point dari program.
// Pertama, kita minta pengguna untuk memasukkan jumlah bilangan genap yang akan dihitung.
// Kemudian, kita panggil fungsi jumlahBilanganGenap dengan jumlah bilangan genap yang dimasukkan oleh pengguna.
// Terakhir, kita tampilkan hasilnya yaitu bilangan genap terakhir dan jumlah total bilangan genap pertama yang telah dihitung.