package main 
import "fmt"

func main() {
	var x,y int

	for {
		fmt.Scan(&x, &y)
		
		terurut(&x, &y)
		if x == y {
			break
		}
		fmt.Println(x,y)
	}
}

func terurut(x, y *int) {
		if *x > *y {
			*x, *y = *y, *x
		}
	}
