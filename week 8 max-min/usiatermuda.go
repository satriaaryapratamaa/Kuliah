package main
import "fmt"

const NMAX int = 256
type pemain struct {
	nama string
	usia, rating int
}

type tabPemain [NMAX]pemain

func bacapemain(A *tabPemain, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i].nama, &A[i].usia, &A[i].rating)
	}
}

func MinUsiaMaxRating(A tabPemain, n int) string {
	var i int
	idx := 0
	i = 1
	for i <= n-1 {
		if A[idx].usia > A[i].usia {
			idx = i
		} else if A[i].usia == A[idx].usia && A[i].rating > A[idx].rating {
			idx = i
		}
		i++
	}
	return A[idx].nama
}

func main() {
	var p tabPemain
	var n int

	fmt.Scan(&n)
	bacapemain(&p,n)
	fmt.Print(MinUsiaMaxRating(p,n))
}