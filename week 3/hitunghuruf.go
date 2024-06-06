package main
import "fmt"

func main() {
	var jumlah int
	var kar byte
	fmt.Scanf("%c", &kar)
	sum(kar, &jumlah)
	fmt.Println(jumlah)
}