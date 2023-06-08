package main

import "fmt"

func main() {
	// 1. Variable
	var name string = "Yudha Islami Sulistya"
	var age int = 20
	var isMarried bool = false
	var weight float32 = 60.5
	var height float64 = 170.5

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Married:", isMarried)
	fmt.Println("Weight:", weight)
	fmt.Println("Height:", height)

	// 2. Operand, Operator and Expression
	x := 10 + 5
	fmt.Println(x)

	a := 10
	b := 5
	c := a + b
	fmt.Println(c)

	// 2.1 Luas Segitiga
	alas := 10
	tinggi := 5
	luas := alas * tinggi / 2
	fmt.Println(luas)

	// 2.2 String Operation
	helloWorld := "Hello" + " " + "World"
	fmt.Println(helloWorld)

	// 4. Structur Control (Branching and Looping)

	// 4.1 Branching If, Else If, Else
	hour := 15
	if hour < 12 {
		fmt.Println("Selamat Pagi")
	} else if hour < 18 {
		fmt.Println("Selamat Sore")
	} else {
		fmt.Println("Selamat Malam")
	}

	// 4.2 Switch
	var today int = 2
	switch today {
	case 1:
		fmt.Println("Senin")
	case 2:
		fmt.Println("Selasa")
	case 3:
		fmt.Println("Rabu")
	case 4:
		fmt.Println("Kamis")
	case 5:
		fmt.Println("Jumat")
	case 6:
		fmt.Println("Sabtu")
	case 7:
		fmt.Println("Minggu")
	default:
		fmt.Println("Hari tidak ditemukan")
	}

	// 4.3 For Loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 4.4 For Loop with Continue and Break
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}

		if i > 3 {
			break
		}
		fmt.Println(i)
	}

	// 4.4 For Loop with Over String
	sentence := "Hello"
	for i := 0; i < len(sentence); i++ {
		fmt.Printf(string(sentence[i]) + "-")
	}
	fmt.Println("    -------->>>")
	for pos, char := range sentence {
		fmt.Printf("Character %c start at byte position %d\n", char, pos)
	}

	// 5. Function
}
