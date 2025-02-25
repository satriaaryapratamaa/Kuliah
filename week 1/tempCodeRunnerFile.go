package main

import "fmt"

func main() {
	var kalimat string
	var i, total int

	fmt.Scan(&kalimat)
	for i = 0; i < len(kalimat); i++ {
		if kalimat[i:i+1] != "a" && kalimat[i:i+1] != "i" && kalimat[i:i+1] != "." {
			total++
		}
	}
	fmt.Println(total)

}
