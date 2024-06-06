package main

import (
	"fmt"
)

const NMAX = 5
type tabINT[NMAX]int

func bacaData(n *int, t *tabINT){
	var bilangan int
	fmt.Scan(&bilangan)
	*n = 0
	for *n <= NMAX-1 && bilangan != -1 {
		t[*n] = bilangan
		*n = *n + 1
		fmt.Scan(&bilangan)
	}
}

func rataRata(n int, t tabINT) float64 {
	var total int
	total = 0 // sebenarnya opsional bisa pakai atau tidak 
	for i := 0; i <= n-1; i++ {
		total = total + t[i]
	}
	return float64(total) / float64(n)
}

func main() {
	var x int
	var rata tabINT

	bacaData(&x, &rata)
	fmt.Println(rataRata(x, rata))
}