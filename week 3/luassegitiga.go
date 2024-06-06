package main
import "fmt"

func main() {
	var n int
	var a, t int
	var l float64

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a, &t)
		hitungluas(a, t, &l)
		fmt.Println(l)
	}
}

func hitungluas(alas, tinggi int, luas *float64) {
	*luas = 0.5 * float64(alas) * float64(tinggi)
}