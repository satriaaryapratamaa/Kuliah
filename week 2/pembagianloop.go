package main

import (
	"fmt"
)

func pembagianLoop(dividen, pembagi int) int {
	var hasilBagi int

	for dividen >= pembagi {
		dividen -= pembagi
		hasilBagi++
	}
	
	return hasilBagi
}

func main () {
	var n, m int
	fmt.Scan(&n, &m)
	fmt.Println(pembagianLoop(n, m))
}