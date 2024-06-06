package main
import (
	"fmt"
	"math"
)
type titik struct {
	x, y float64
}

func main() {
	var A, B, C titik
	var dab, dac float64
	
	fmt.Scan(&A.x, &A.y)
	fmt.Scan(&B.x, &B.y)
	fmt.Scan(&C.x, &C.y)
	
	dab = math.Sqrt(math.Pow((A.x - B.x),2) + math.Pow((A.y - B.y),2))
	dac = math.Sqrt(math.Pow((A.x - C.x),2) + math.Pow((A.y - C.y),2))
	fmt.Println(dac < dab)
	
}