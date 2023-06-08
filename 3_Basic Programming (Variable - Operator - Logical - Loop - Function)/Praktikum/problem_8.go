package main

import "fmt"

func cetakTablePerkalian(number int) {
	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {
			fmt.Print(i * j)
			if j < number {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	cetakTablePerkalian(9)
}
