package main

import "fmt"

type pemain struct {
	nama        string
	gol, assist int
}

func main() {
	var a, b, c pemain
	//var golTerbesar, assistTerbesar string

	fmt.Scan(&a.nama, &a.gol, &a.assist)
	fmt.Scan(&b.nama, &b.gol, &b.assist)
	fmt.Scan(&c.nama, &c.gol, &c.assist)

	if a.gol > b.gol {
		fmt.Println(a.nama)
	} else if b.gol > a.gol {
		fmt.Println(b.nama)
	} else {
		fmt.Println(c.nama)
	}

	if a.assist > b.assist {
		fmt.Println(a.nama)
	} else if b.assist > a.assist {
		fmt.Println(b.nama)
	} else {
		fmt.Println(c.nama)
	}

}
