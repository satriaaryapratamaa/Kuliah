package main
import "fmt"

func baris(x, y int) {

	for i := x; i <= y; i++ {
		fmt.Print(i, " ")
	}
}

func main() {
	var m, n int
	
	fmt.Scan(&m, &n)
	baris(m, n)
}
