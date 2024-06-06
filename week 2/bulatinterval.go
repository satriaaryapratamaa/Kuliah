// Buatlah fungsi jumlah untuk menjumlahkan bilangan bulat dalam interval tertentu. Misalnya, 
//bila fungsi diminta untuk menjumlahkan bilangan dalam interval 2 hingga 5, maka fungsi akan mengembalikan 14, 
//yaitu hasil penjumlahan dari 2, 3, 4, dan 5. Bila fungsi diminta untuk menjumlahkan bilangan dalam interval 10 hingga 6, 
//maka fungsi akan mengembalikan 40, yaitu hasil penjumlahan dari 10, 9, 8, 7, dan 6 . 

// Fungsi dipanggil dalam program dengan masukan dan keluaran sebagai berikut:

// Masukan berupa dua buah bilangan bulat yang menyatakan interval.
// Keluaran adalah hasil penjumlahan bilangan dalam interval masukan.

package main 
import "fmt"

func jumlah(x, z int) int {
	var y int

	if x < z {
		for i := x; i <= z; i++ {
			y += i
		}
		return y
	} 
	
	if z < x {
		for j := z; j <= x; j++ {
			y += j
		}
		return y
	}
	return y
}

func main() {
	var bil1, bil2 int
	fmt.Scan(&bil1, &bil2)
	fmt.Println(jumlah(bil1, bil2))
}