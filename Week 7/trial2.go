package main
import "fmt"
type arrReal[5]float64

func main() {
	var tabReal arrReal
	var nilai float64
	var i int

	i = 0
	fmt.Scan(&nilai)
	for i < 5 && nilai > 0 {
		tabReal[i] = nilai
		fmt.Scan(&nilai)
		i = i + 1
	} 
	fmt.Println(tabReal)
}
