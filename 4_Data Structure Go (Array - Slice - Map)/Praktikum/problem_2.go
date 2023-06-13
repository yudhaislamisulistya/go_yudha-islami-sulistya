package main

import (
	"fmt"
	"strconv"
)

// Sample Test Case
// Input: 76523752
// Output: [6 3]
// Input: 1122
// Output: []

func munculSekali(angka string) []int {
	var arrayAngka []int
	for _, value := range angka {
		angka, _ := strconv.Atoi(string(value))
		arrayAngka = append(arrayAngka, angka)
	}
	var arrayAngkaNew []int
	for index, value := range arrayAngka {
		var duplicate bool = false
		for index2, value2 := range arrayAngka {
			if value == value2 && index != index2 {
				duplicate = true
			}
		}
		if !duplicate {
			arrayAngkaNew = append(arrayAngkaNew, value)
		}
	}
	return arrayAngkaNew
}

func main() {
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}
