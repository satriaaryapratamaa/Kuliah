package main

import "fmt"

func cekfaktorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * cekfaktorial(n-1)
	}
}

func main() {
	var x int = 5
	fmt.Println(cekfaktorial(x))
}
