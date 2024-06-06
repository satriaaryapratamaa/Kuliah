package main
import "fmt"

func main() {
	var c, r,f,k float64

	fmt.Scan(&c)
	konversiCelcius(c, &r, &f, &k)
	fmt.Print(r,"R ", f, "F ", k, "K")

}

func konversiCelcius (celcius float64, reamur, fahrenheit, kelvin *float64) {
	*reamur = 5/4 *celcius
	*fahrenheit = (celcius * 9/5) + 32
	*kelvin = celcius + 273.5

}