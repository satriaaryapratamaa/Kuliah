//SOAL

// Di sebuah planet bernama Cerulia terdapat beberapa koin yang digunakan sebagai alat tukar, yaitu Quantum Shard, Galactum, Nebula Crown, Stellaris, dan Lumin.
// Satu koin Quantum Shard setara dengan 10 Galactum. Satu koin Galactum setara dengan 3 Nebula Crown.
// Satu koin Nebula Crown setara dengan 2 Stellaris. Satu koin Stellaris setara dengan 20 Lumin.

// Masukan berupa uang dalam koin Lumin dengan tipe bilangan bulat.

// Keluaran berupa uang dalam koin Quantum Shard, Galactum, Nebula Crown, Stellaris, dan Lumin.

package main
import "fmt"

func main() {
	var (
		lumin     int
		quantum   int
		galactum  int
		nebula    int
		stellaris int
	)

	fmt.Scan(&lumin)

	quantum = lumin / 1200
	lumin = lumin % 1200

	galactum = lumin / 120
	lumin = lumin % 120

	nebula = lumin / 40
	lumin = lumin % 40

	stellaris = lumin / 20
	lumin= lumin % 20


	fmt.Print(quantum, galactum, nebula, stellaris,lumin)

}
