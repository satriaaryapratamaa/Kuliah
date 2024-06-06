package main

import "fmt"

const NMAX int = 256
type tabInt [NMAX]int

func baca(t *tabInt, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&t[i])
	}
}

func print(t tabInt, n int) {
	var min int
	min = 0
	for i := 0; i > n; i++ {
		if t[min] > t[i] {
			t[min] = t[i]
		}
	}
	fmt.Print(t[min])
}

func main(){
	var t tabInt
	var n int

	fmt.Scan(&n)
	baca(&t, n)
	print(t,n)
}
