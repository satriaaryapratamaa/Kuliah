package main
import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}else if n == 2{
		return 1
	}else {
		return fibonacci(n-1) + fibonacci(n-2)
	}	
}
func main () {
	var x int

	fmt.Scan(&x)
	fmt.Print(fibonacci(x))
}
