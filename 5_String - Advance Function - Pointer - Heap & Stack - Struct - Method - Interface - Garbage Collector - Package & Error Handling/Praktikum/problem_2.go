package main

import "fmt"

func Caesar(offset int, str string) string {
	result := ""
	offset = offset % 26 // Memastikan offset dalam rentang 0-25

	for i := 0; i < len(str); i++ {
		char := str[i]

		// Cek karakter alfabet besar
		if char >= 'A' && char <= 'Z' {
			ascii := int(char) - 'A'
			shifted := (ascii + offset) % 26
			result += string(shifted + 'A')
		} else if char >= 'a' && char <= 'z' { // Cek karakter alfabet kecil
			ascii := int(char) - 'a'
			shifted := (ascii + offset) % 26
			result += string(shifted + 'a')
		} else { // Karakter selain alfabet tidak diubah
			result += string(char)
		}
	}

	return result
}

func main() {
	fmt.Println(Caesar(3, "abc"))
	fmt.Println(Caesar(2, "alta"))
	fmt.Println(Caesar(10, "alterraacademy"))
	fmt.Println(Caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(Caesar(1000, "abcdefghijklmnopqrstuvwxyz"))
}
