package main

import "fmt"

func main() {
	// Sample Test Cases
	// Input: 6
	// Output 1 2 3 6
	// Input: 20
	// Output 1 2 4 5 10 20

	var bilangan int
	fmt.Scanf("%d", &bilangan)

	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Println(i)
		}
	}
}
