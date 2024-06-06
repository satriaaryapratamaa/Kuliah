package main

import "fmt"

func main() {
	type liga struct {
		nama string
		poin int
		selisihgol int
	}

	var klubA, klubB, klubC liga
	
		fmt.Scan(&klubA.nama)
		fmt.Scan(&klubA.poin)
		fmt.Scan(&klubA.selisihgol)
		fmt.Println("")

		fmt.Scan(&klubB.nama)
		fmt.Scan(&klubB.poin)
		fmt.Scan(&klubB.selisihgol)
		fmt.Println("")

		fmt.Scan(&klubC.nama)
		fmt.Scan(&klubC.poin)
		fmt.Scan(&klubC.selisihgol)
		fmt.Println("")

		ratarata := float64(klubA.poin + klubB.poin + klubC.poin) / 3
		fmt.Print(ratarata)
}

