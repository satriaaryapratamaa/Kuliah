package main
import "fmt"

const NMAX = 100

type tim struct {
	a, b int
}
type tabGol [NMAX]tim

func input(t *tabGol, i *int){
	//fmt.Scan(i)
	for i < NMAX {
		for t[*i].a > 0 && t[*i].b > 0{
		fmt.Scan(&t[*i].a, &t[*i].b) 
		*i++
		}
	}
}

func output(t tabGol, n int) (float64, float64) {
	var a, b int
	var hasilA, hasilB float64
	for i := 0; i < n; i++ {
		a += t[i].a
		b += t[i].b 
		
	}
	hasilA = float64(a) / float64(n)
	hasilB = float64(b) / float64(n)
	return hasilA, hasilB
}

func main() {	
	var t tabGol
	var n int
	fmt.Scan(&n)
	input(&t, &n)
	fmt.Print(output(t,n))

}