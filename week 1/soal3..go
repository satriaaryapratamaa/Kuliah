package main

import "fmt"

func main() {
	var (
		x bool
		y bool
	)

	fmt.Scan(&x, &y)

	if x == true || y == true {
		fmt.Print("berlibur")
	} else {
		fmt.Print("tida berlibur")
	}
}