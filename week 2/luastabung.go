package main
import "fmt"

func main() {
	var jari2, tinggi float64

	fmt.Scan(&jari2, &tinggi)
	fmt.Printf("%.3f\n", luastabung(jari2, tinggi))
}

func luastabung (jari2, tinggi float64) float64 {
	const phi = 3.14
	return 2 * phi * jari2 * (jari2 + tinggi)
}

