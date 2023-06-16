package main

import "fmt"

// pow with time complexity O(log n) or logarithmic time complexity
func pow(x, n int) int {
	if n == 0 {
		return 1
	}
	if n%2 == 1 {
		return x * pow(x, n-1)
	}
	return pow(x*x, n/2)
}

func main() {
	fmt.Println(pow(2, 3))
	fmt.Println(pow(5, 3))
	fmt.Println(pow(10, 2))
	fmt.Println(pow(2, 5))
	fmt.Println(pow(7, 3))
}
