package main

import "fmt"

func main() {
	var (
		jam int
		menit int
		detik int
	)
	fmt.Scan(&jam, &menit, &detik)
	jam = jam * 3600
	menit = menit * 60
	detik = detik 

	hasil := jam + menit + detik

	fmt.Println(hasil)
}

