// Buatlah fungsi (dalam bentuk pseudocode) untuk menghitung diskon yang 
// berdasarkan pada total belanja, dan status membership. Jika total belanja > 
// 100.000 dan member, maka diskon 10%, jika belanja > 100.000 dan non-member, 
// maka diskon 5%. 
// Selanjutnya buatlah program yang meminta masukan dari pengguna, memanggil 
// fungsi, menghitung total belanja akhir dan menampilkan total belanja akhir 
// tersebut. 

// Masukan terdiri dari dua nilai, yaitu bilangan bulat menyatakan total belanja awal 
// dan boolean yang menyatakan status membership. 

// Keluaran adalah bilangan yang menyatakan total belanja akhir setelah dipotong 
// diskon apabila memperoleh diskon. 

package main 
import "fmt"

func diskon(x int, y bool) int {

	if x >= 100000 && y == true {
		return x * 90/100
	} else if x >= 100000 && y == false{
		return x * 95/100
	} else {
		return x
	}
	
}

func main() {
	var total, z int
	var cekDiskon bool

	fmt.Scan(&total, &cekDiskon)
	z = diskon(total,cekDiskon)
	fmt.Println(z)
}

////>>>PESUTKOD

// function diskon(total : integer, cekDiskon : boolean) -> integer
// {mengembalikan hasil dari total * 90/100 jika total >= 100000 dan total * 95/100 jika total >= 100000}
//  kamus
// 		total : integer
// 		cekDiskon : boolean
// 	algoritma
// 		if total >= 100000 and cekDiskon == true then
// 			return total * 90/100
// 		else if total >= 100000 && cekDiskon == false then
// 			return total * 95/100
// 		else 
// 			return total
// 		endif
// endfunction

// program diskon
// 	kamus 
// 		total, z : integer
// 		cekDiskon : boolean
// 	algoritma
// 		input(total,cekDiskon)
//		z <- diskon(total,cekDiskon)
// 		output(z)
// endprogram