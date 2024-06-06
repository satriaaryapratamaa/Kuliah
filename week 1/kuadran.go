// SOAL
// Layar navigasi X-Wing milik pasukan pemberontak dapat menentukan posisi pesawat musuh berada di kuadran mana.
// Buatlah program untuk menentukan posisi kuadran pesawat musuh dalam koordinat 2D kartesius dengan syarat berikut:

// kuadran I, jika x positif dan y positif.
// kuadran II, jika x negatif dan y positif.
// kuadran III, jika x negatif dan y negatif.
// kuadran IV, jika x positif dan y negatif.
// Masukan berupa dua bilangan real yang menyatakan posisi pesawat dalam x dan y.

// Keluaran berupa string yang menyatakan nama kuadran.

package main 
import "fmt"

func main() {
	var (
		x float64
		y float64
	)

	fmt.Scan(&x, &y)

	if x > 0 && y > 0 {
		fmt.Println("Kuadran 1")
	} 
	if x < 0 && y > 0 {
		fmt.Println("Kuadran 2")	
	} 
	if x < 0 && y < 0 {
		fmt.Println("Kuadran 3")
	} 
	if x > 0 && y < 0 {
		fmt.Println("Kuadran 4")
	}

}