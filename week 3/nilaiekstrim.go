package main
import "fmt"

func cariEkstrim(n1, n2 int, min, max *int) {
	if n1 < n2 {
		*min = n1
		*max = n2
	} else {
		*min = n2
		*max = n1
	}
}

func main() {
	var a,b,c,d int
	var min1,min2,max1,max2,min,max,temp int

	fmt.Scan(&a, &b, &c, &d)
	cariEkstrim(a, b, &min1, &max1)
	cariEkstrim(c, d, &min2, &max2)
	cariEkstrim(max1, max2, &temp, &max)
	cariEkstrim(min1, min2, &min, &temp)
	fmt.Println(max, min)
}

