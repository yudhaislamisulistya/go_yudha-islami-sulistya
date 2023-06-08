package main

import "fmt"

func main() {
	// Sample Test cases
	// Input: T = 20, r = 3
	// Output: 602.88
	// Menghitung Luas Permukaan Tabung

	T := 20
	r := 4

	const pi float64 = 3.14

	luas_tutup := pi * float64(r) * float64(r)
	luas_selimut := 2 * pi * float64(r) * float64(T)
	luat_alas := pi * float64(r) * float64(r)

	luas_permukaan := luas_tutup + luas_selimut + luat_alas

	fmt.Println(luas_permukaan)

}
