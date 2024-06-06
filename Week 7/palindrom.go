package main
import "fmt"

const NMAX = 256
type nBilangan[NMAX]int

func isiArray(t *nBilangan, x *int){
	fmt.Scan(x)
	for i := 0; i < *x; i++{
		fmt.Scan(&t[i])
	}
}

func reverseArray(t1 nBilangan, x int, t2 *nBilangan) {
	var j int
	j = x - 1
	for i := 0; i < x; i++{
		t2[i] = t1[j]
		j = j - 1
	}
}

func cekPalindrom(t nBilangan, x int) bool {
	var temp nBilangan
	var status bool = true
	reverseArray(t, x, &temp)
	for i := 0; i < x && status; i++ {
		status = status && (t[i] == temp[i])
	}
	return status
}

func main() {
	var A nBilangan
	var x int

	isiArray(&A, &x)
	fmt.Println(cekPalindrom(A,x))
}