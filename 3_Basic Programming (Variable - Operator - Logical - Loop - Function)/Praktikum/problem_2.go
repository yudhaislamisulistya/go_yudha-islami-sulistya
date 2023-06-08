package main

import "fmt"

func main() {
	// Grade Nilai
	// Nilai 80 - 100: 1
	// Nilai 65 - 79: 2
	// Nilai 50 - 64: 3
	// Nilai 35 - 49: 4
	// Nilai 0 - 34: 5
	// Nilai Kurang dari 0 atau lbeih dari 100 maka tampilan "Nilai Invalid"

	var studentScore int = 80

	if studentScore >= 80 && studentScore <= 100 {
		fmt.Println("A")
	} else if studentScore >= 65 && studentScore <= 79 {
		fmt.Println("B")
	} else if studentScore >= 50 && studentScore <= 64 {
		fmt.Println("C")
	} else if studentScore >= 35 && studentScore <= 49 {
		fmt.Println("D")
	} else if studentScore >= 0 && studentScore <= 34 {
		fmt.Println("E")
	} else {
		fmt.Println("Nilai Invalid")
	}
}
