package main
import "fmt"

func displayFactors(N int, N2 *int) {
	if *N2 > N {
		return
	}
	if N % *N2 == 0 {
		fmt.Print(*N2, " ")
	}
	*N2++
	displayFactors(N, N2)
}

func main() {
	var N, N2 int
	N2 = 1
	fmt.Scanln(&N)
	displayFactors(N, &N2)
	fmt.Println()
}