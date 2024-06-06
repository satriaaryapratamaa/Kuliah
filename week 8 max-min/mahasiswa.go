// Satria Arya Pratama
// IF-47-11 (103012300250)

package main

import (
	"fmt"
)

type Wisudawan struct {
	Nama     string
	NIM      string
	Eprt     int
	Semester int
	Ipk      float64
}

// inputData mengisi array wisudawan sampai NIM "none" ditemukan
func inputData(wisudawanArr *[]Wisudawan) {
	for {
		var input Wisudawan
		fmt.Scan(&input.Nama, &input.NIM, &input.Eprt, &input.Semester, &input.Ipk)
		*wisudawanArr = append(*wisudawanArr, input)
		if input.NIM == "none" {
			break
		}
	}
}
// eprtTertinggi mencari nilai Eprt tertinggi dari array wisudawan
func eprtTertinggi(wisudawanArr []Wisudawan) int {
	maxEprt := 0
	for _, wisudawan := range wisudawanArr {
		if wisudawan.Eprt > maxEprt {
			maxEprt = wisudawan.Eprt
		}
	}
	return maxEprt
}

// ipkTerendah mencari nilai Ipk terendah dari array wisudawan
func ipkTerendah(wisudawanArr []Wisudawan) float64 {
	if len(wisudawanArr) == 0 {
		return 0
	}
	minIpk := wisudawanArr[0].Ipk
	for _, wisudawan := range wisudawanArr {
		if wisudawan.Ipk < minIpk {
			minIpk = wisudawan.Ipk
		}
	}
	return minIpk
}

// rataRataSemester menghitung rata-rata semester dari array wisudawan
func rataRataSemester(wisudawanArr []Wisudawan) float64 {
	totalSemester := 0
	for _, wisudawan := range wisudawanArr {
		totalSemester += wisudawan.Semester
	}
	if len(wisudawanArr) == 0 {
		return 0
	}
	return float64(totalSemester) / float64(len(wisudawanArr))
}

func main() {
	var wisudawanArr []Wisudawan
	wisudawanArr = make([]Wisudawan, 0, 1000)

	inputData(&wisudawanArr)

	maxEprt := eprtTertinggi(wisudawanArr)
	minIpk := ipkTerendah(wisudawanArr)
	avgSemester := rataRataSemester(wisudawanArr)

	fmt.Println("tertinggi eprt:",maxEprt, "terendah ipk:", minIpk, "rata rata semester:", avgSemester)
}
