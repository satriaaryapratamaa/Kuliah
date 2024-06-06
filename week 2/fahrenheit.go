package main
import "fmt"

func main() {
	var x int
	var c float64
	
	fmt.Scan(&x)
	for i := 0; i < x; i++ {
		fmt.Scan(&c)
		fmt.Println(fahrenheit(c))
	}
}

func fahrenheit(celcius float64) float64 {
	return celcius * 9/5 + 32
}
