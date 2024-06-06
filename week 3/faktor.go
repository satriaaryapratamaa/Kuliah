// Buatlah prosedur cekFaktor untuk menentukan apakah bilangan bulat a merupakan faktor dari bilangan b atau bukan.

// Prosedur dipanggil dalam program dengan masukan (in) dan keluaran sebagai berikut:

// Masukan berupa dua buah bilangan bulat positif, a dan b.
// Keluaran mencetak "Ya" jika a adalah faktor dari b atau "Tidak" jika bukan.

package main
import "fmt"

func main() {
	var a, b int

	fmt.Scan(&a, &b)
	faktor(&a, &b)
}

func faktor(a, b *int) {
	
	if *b % *a == 0{
		fmt.Print("Ya, ", *a, " faktor ", *b)
	} else {
		fmt.Print("Tidak, ", *a, " bukan faktor ", *b)
	}
}
