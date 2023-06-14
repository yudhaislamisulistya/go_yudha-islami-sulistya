package main

import (
	"fmt"
	"strings"
)

func Compare(A string, B string) string {
	// if contains A in B
	if strings.Contains(B, A) {
		return A
	} else {
		return B
	}
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANGOORO", "KANG"))
	fmt.Println(Compare("KI", "KIJANG"))
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))
	fmt.Println(Compare("ILALANG", "ILA"))
}
