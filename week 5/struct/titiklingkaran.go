package main 
import (
	"fmt"
	"math"
)

func main() {
	type titik struct {
		x , y float64
	}
	var a,b,c titik
	var dAB, dAC float64

	fmt.Scan(&a.x, &a.y)
	fmt.Scan(&b.x, &b.y)
	fmt.Scan(&c.x, &c.y)

	dAB = math.Sqrt(math.Pow((a.x - b.x),2) + math.Pow((a.y - b.y),2))
	dAC = math.Sqrt(math.Pow((a.x - c.x),2) + math.Pow((a.y - c.y),2))

	fmt.Println(dAC <= dAB)
}

