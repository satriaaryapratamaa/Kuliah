package main
import "fmt"

type balok struct {
	panjang, lebar, tinggi float64
}

func luasBalok(x balok) float64 {
	var luas float64
	fmt.Print("Luas balok = ")
	luas = 2 * ((x.panjang * x.lebar)+ (x.panjang * x.tinggi) + (x.lebar * x.tinggi))
	return luas
}

func volumeBalok(x balok) float64 {
	var volume float64
	fmt.Print("Volume balok = ")
	volume = x.panjang * x.lebar * x.tinggi
	return volume
}

// func bacaLuasVolume(x balok) {

// }

func main() {
	var bb balok

	fmt.Scan(&bb.panjang, &bb.lebar, &bb.tinggi)
	fmt.Println(luasBalok(bb))
	fmt.Println(volumeBalok(bb))
	// bacaLuasVolume(bb)
}