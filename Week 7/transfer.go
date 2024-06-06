package main
import "fmt"

const NMAX = 256
type dataTransfer[NMAX]struct {
	nama, jenisTrans string
	hargaTrans int
}

func hitungTrans(x dataTransfer, y *int) {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&x[i].jenisTrans, &x[i].nama, &x[i].hargaTrans)
		*y++
	}
}

func bacaTransfer(x dataTransfer, y int) {
	var total int
	total = 0
	for i := 0; i < y; i++ {
		if x[i].jenisTrans == "in"{
			total -= x[i].hargaTrans
		} else if x[i].jenisTrans == "out" {
			total += x[i].hargaTrans
		}
	}
	fmt.Print(total > 0)
}

func main() {
	var n dataTransfer
	var y int

	fmt.Scan(&n)
	hitungTrans(n, &y)
	bacaTransfer(n, y)
}