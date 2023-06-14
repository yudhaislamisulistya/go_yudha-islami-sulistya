package main

import "fmt"

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	nameEncode := ""
	mapping := map[string]string{
		"r": "i",
		"i": "r",
		"z": "a",
		"k": "p",
		"y": "b",
	}
	for i := 0; i < len(s.name); i++ {
		nameEncode += mapping[string(s.name[i])]
	}
	return nameEncode
}

func (s *student) Decode() string {
	nameDecode := ""
	mapping := map[string]string{
		"i": "r",
		"r": "i",
		"a": "z",
		"p": "k",
		"b": "y",
	}
	for i := 0; i < len(s.nameEncode); i++ {
		nameDecode += mapping[string(s.nameEncode[i])]
	}
	return nameDecode
}

func main() {
	var menu int
	var a = student{}
	var c Chiper = &a
	fmt.Print("[1] Encrypt\n[2] Decrypt\nChoose: ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("Input Name: ")
		fmt.Scan(&a.name)
		fmt.Println("Name Encode:", a.name, "->", c.Encode())
	} else if menu == 2 {
		fmt.Print("Input Name Encode: ")
		fmt.Scan(&a.nameEncode)
		fmt.Println("Name Decode:", a.nameEncode, "->", c.Decode())
	} else {
		fmt.Println("Menu Not Found")
	}
}
