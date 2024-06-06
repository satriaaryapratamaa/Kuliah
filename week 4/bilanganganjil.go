package main

import "fmt"

func main() {
	var x int 
	fmt.Scan(&x)
	ganjil(x)
}

func ganjil(n int) {
	if n == 1 {
		fmt.Print(1)
	} else {
		if n%2 == 0 {
			ganjil(n - 1)
		} else {
			ganjil(n - 2)
			fmt.Print(n)
		}
	}
}
