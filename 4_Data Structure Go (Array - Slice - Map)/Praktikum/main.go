package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 1. Array
	var primes [5]int
	var countries [5]string

	fmt.Println(reflect.ValueOf(primes).Kind())
	fmt.Println(reflect.ValueOf(countries).Kind())

	// 1.1 Array Spasial
	x := [5]int{1, 3, 5, 7, 9}
	var y [5]int = [5]int{1, 3, 5}
	fmt.Println(x)
	fmt.Println(y)

	// 1.2 Array Looping
	primes_new := [5]int{1, 3, 5, 7, 9}
	// 1.2.1 Array Looping Cara Pertama
	for index := 0; index < len(primes_new); index++ {
		fmt.Println(primes_new[index])
	}
	// 1.2.2 Array Looping Cara Kedua
	for index, element := range primes_new {
		fmt.Println(index, element)
	}
	// 1.2.3 Array Looping Cara Ketiga
	index := 0
	for index < len(primes_new) {
		fmt.Println(primes_new[index])
		index++
	}
	// 1.2.4 Array Looping Cara Keempat
	index_new := 0
	for range primes {
		fmt.Println(primes_new[index_new])
		index_new++
	}

	// 2. Slice
	slice_a := []int{1, 3, 5, 7, 9}
	fmt.Println("Len Slice A: ", len(slice_a))
	fmt.Println("Cap Slice A: ", cap(slice_a))
	slice_a = append(slice_a, 11)
	fmt.Println("Len Slice A: ", len(slice_a))
	fmt.Println("Cap Slice A: ", cap(slice_a))
	slice_a = append(slice_a, 13)
	fmt.Println("Len Slice A: ", len(slice_a))
	fmt.Println("Cap Slice A: ", cap(slice_a))
	fmt.Println(reflect.ValueOf(slice_a).Kind())

	var slice_b = make([]int, 5, 10)
	fmt.Println("Len Slice B: ", len(slice_b))
	fmt.Println("Cap Slice B: ", cap(slice_b))

	// 2.1 Slice Append and Copy
	var colors = []string{"red", "green", "yellow"}
	colors = append(colors, "blue")
	fmt.Println(colors)
	var colors_copy = make([]string, len(colors))
	copy(colors_copy, colors)
	fmt.Println(colors_copy)

	// 2.2 Slice Looping
	var slice_c = []int{1, 3, 5, 7, 9}
	for index, element := range slice_c {
		fmt.Println(index, element)
	}

	// 2.2 Slice Pointer
	var slice_d = []int{1, 3, 5, 7, 9}
	var slice_e = slice_d
	slice_e[0] = 100
	fmt.Println(slice_d)
	fmt.Println(slice_e)

	// 3. Map
	var map_a = map[string]int{"foo": 1, "bar": 2}
	fmt.Println(map_a)

	// 3.1 Map Append and Copy
	var map_b = map[string]int{"foo": 1, "bar": 2}
	map_b["baz"] = 3
	fmt.Println(map_b)
	var map_c = make(map[string]int)
	for key, value := range map_b {
		map_c[key] = value
	}
	fmt.Println(map_c)
}
