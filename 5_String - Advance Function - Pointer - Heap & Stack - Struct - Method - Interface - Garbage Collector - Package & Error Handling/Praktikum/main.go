package main

import (
	"fmt"
	"strings"
)

const (
	str    = "something"
	substr = "some"
)

// 4. Struct
type User struct {
	FirstName, LastName string
}

// 5. Method
type Employee struct {
	Name string
	Age  int
}

func (employee Employee) sayHello() {
	fmt.Println("Hello", employee.Name)
}

func main() {
	// 1. String
	// 1.1 Len String
	sentence := "Hello World"
	lenSentence := len(sentence)
	fmt.Println(lenSentence)
	// 1.2 Compare String
	str1 := "Hello"
	str2 := "Hello"
	if str1 == str2 {
		fmt.Println("Sama")
	} else {
		fmt.Println("Tidak Sama")
	}
	// 1.3. Contain String
	res := strings.Contains(str, substr)
	fmt.Println(res)
	// 1.4. Show String by Index
	value := "cat;dog"
	substring := value[4:]
	fmt.Println(substring)
	// 1.5. Replace
	s := "Katak"
	t := strings.Replace(s, "a", "o", 1)
	fmt.Println(t)
	// 1.6. Insert
	p := "Ktk"
	index := 2
	q := p[:index] + "a" + p[index:]
	fmt.Println(q)

	// 2. Advance Function
	// 2.1. Function as Parameter
	fmt.Println(sum(12, 10))
	// 2.2. Anonymous Function or Literal Function
	bagianawal := func() {
		fmt.Println("Hello World")
	}
	bagianawal()
	bagianKedua := func(urutan string) {
		fmt.Println("Hello World " + urutan)
	}
	bagianKedua("Kedua")
	// 2.3. Closure Function
	bagianKetiga := func(urutan string) func() {
		return func() {
			fmt.Println("Hello World " + urutan)
		}
	}
	bagianKetiga("Ketiga")()

	counter := newCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	// 2.4 Defer Function
	defer func() {
		fmt.Println("Later")
	}()
	fmt.Println("First")
	// 3. Pointer
	name := "John"
	pointer := &name
	*pointer = "Doe"
	fmt.Println(name)
	fmt.Println(pointer)
	fmt.Println(*pointer)
	// 4. Struct
	var user User
	user.FirstName = "John"
	user.LastName = "Doe"
	fmt.Println(user)

	user1 := User{
		FirstName: "John",
		LastName:  "Doe",
	}
	fmt.Println(user1)
	// 5. Method
	employee := Employee{
		Name: "John",
		Age:  21,
	}
	employee.sayHello()

}

// 2. Advance Function
func sum(data ...int) int {
	result := 0
	for _, value := range data {
		result += value
	}
	return result
}

func newCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
