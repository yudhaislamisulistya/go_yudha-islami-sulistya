package main

import "fmt"

// Sample Test Cases
// Input: ['kazuya', 'jin', 'lee'], ['kazuya', 'feng']
// Output: ['kazuya', 'jin', 'lee', 'feng']
// Input: ['lee', 'jin'], ['kazuya', 'panda']
// Output: ['lee', 'jin', 'kazuya', 'panda']

func ArrayMerge(arrayA, arrayB []string) []string {
	arrayC := append(arrayA, arrayB...)
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range arrayC {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	arrayC = list
	return arrayC
}

func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))

	fmt.Println(ArrayMerge([]string{}, []string{}))
}
