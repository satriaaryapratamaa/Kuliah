package main

import "fmt"

type arrReal [5]float64

func main() {
	var tabReal arrReal
	var elemen, i int

	i = 0
	fmt.Scan(&elemen)
	for i < elemen && i < 5 {
		fmt.Scan(&tabReal[i])
		i = i + 1
	}
	fmt.Print(tabReal)
}
