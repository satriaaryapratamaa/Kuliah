package main

import "fmt"

const maxTeman = 100
type teman [maxTeman]string

func tulis(m int, t *teman) {
	for i := 0; i < m; i++ {
		fmt.Scan(&t[i])
	}
}

func baca(m int, t teman) {
	for i := 0; i < m; i++ {
		fmt.Print(t[i])
		fmt.Print(" ")
	}
}

func main() {
	var t teman
	var n int

	fmt.Scan(&n)
	tulis(n, &t)
	baca(n, t)
}
