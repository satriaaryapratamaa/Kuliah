package main 

import "fmt"

func main() {
	var n, m int
	var jumlahDeret float64

	fmt.Scan(&n, &m)

	for i := n; i <= m; i++ {
		jumlahDeret = jumlahDeret + 4 / float64(i)
	}
	fmt.Printf("%.2f\n", jumlahDeret)
}

