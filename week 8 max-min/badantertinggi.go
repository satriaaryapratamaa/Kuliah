package main
import "fmt"

const NMAX int = 256
type student struct {
	nama string
	age, height int
}
type tabStudent [NMAX]student

func bacaStudent(t *tabStudent, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&t[i].nama, &t[i].age, &t[i].height)
	}
}

/*
func printStudent(t tabStudent, n int) string {
	var i int
	maxIndex := 0
	for maxIndex < n-1 && t[maxIndex].age != 18 {
		maxIndex++
	}
	i = maxIndex + 1

	for i <= n-1 {
		if t[i].height > t[maxIndex].height && t[i].age == 18 {
			maxIndex = i
		}
		i++
	}
	return t[maxIndex].nama
}
*/

func printStudent(t tabStudent, n int) string {
	var idx, i int
	idx = 0
	for idx < n-1 && t[i].age == 18 {
		i++
	}
	i = idx + 1

	for i < n-1 {
		if t[i].height > t[idx].height && t[i].age == 18 {
			idx = i 
		}
		i++
	}
	return t[idx].nama
} 

func main() {
	var students tabStudent
	var n int

	fmt.Scan(&n)
	bacaStudent(&students, n)
	fmt.Println(printStudent(students, n)) // Print the result to standard output
}