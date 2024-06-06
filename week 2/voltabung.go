package main
import "fmt"

func main() {
	//const phi = 3.14
	var jari2, tinggi float64

	fmt.Scan(&jari2, &tinggi)
	fmt.Printf("%.1f\n", voltabung(jari2, tinggi))

}

func voltabung(ja, t float64) float64 {
	const phi = 3.14
	var vol float64

	vol = phi * (ja*ja) * t
	return vol
}

