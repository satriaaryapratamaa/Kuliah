package main
import "fmt"

const NMAX = 256
type tabINT [NMAX]int

func baca(x int, t *tabINT){
	//var total int
	for i := 0; i < x; i++{
		fmt.Scan(&t[i])
	}
}

func cetak(x int, t tabINT){
	var total int
	for i := 0; i < x; i++ {
		total += t[i]
	}
	fmt.Println(total)
}

func main() {
	var n tabINT
	var x int

	fmt.Scan(&x)
	baca(x, &n)
	cetak(x, n)
	
}