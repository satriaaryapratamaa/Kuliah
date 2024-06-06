package main
import "fmt"

func permutasi(a, c int) int {
	var p, nr int
	
	nr = a - c
	if a >= c {
		for i := 1; i <= a; i-- {
			p *= a
		}
		for j := 1; j <= nr; j--  {
			nr *= c
		}
	}
	return p/nr
}

func main () {
	var a, c int
	fmt.Scan(&a, &c)
	//fmt.Scan(&b, &d)
	
	fmt.Println(permutasi(a, c))

}