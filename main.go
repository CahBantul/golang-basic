package main

import "fmt"

func main() {
	// fmt.Println("Hello World")
	// var firstName string = "Fardan "
	// var lastName string = "Nozami"
	// var name = firstName + lastName
	// fmt.Println(len(name))
	// var isBoolean bool = false
	// fmt.Println(isBoolean)
	// fmt.Println(!false)
	// var int8 int = 32900
	// fmt.Println(int8)
	// var a = 4
	// b := 5
	// fmt.Println(a + b)
	// fmt.Println(a * b)
	var numbers []string
	numbers = append(numbers, "satu", "dua", "tiga")
	fmt.Println(numbers[0:2])

	person := make(map[string]string)
	person["name"] = "ajitama"
	person["job"] = "developer"
	fmt.Println(person)

	students := map[string]int{
		"Nozami": 90,
		"Fardan": 70,
	}

	fmt.Println(students)
}
