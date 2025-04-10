package main

import "fmt"

func sayHello(name string) {
	fmt.Println("hallo", name)
}

func tambah(a, b int) int {
	return a + b
}

func calculate(a, b int) (sum int, multiply int) {
	sum = a + b
	multiply = a * b
	return
}

func main() {
	sayHello("Ajitama")
	result := tambah(5, 10)
	fmt.Println("hasil pemjumlahan", result)
	add, multiply := calculate(3, 6)
	fmt.Println("hasil pemjumlahan", add)
	fmt.Println("hasil perkalian", multiply)

	// if else
	// angka := 9

	// if angka%2 == 0 {
	// 	fmt.Println("Genap")
	// } else {
	// 	fmt.Println("Ganjil")
	// }

	// multiple if
	// nilai := 90

	// if nilai >= 90 {
	// 	fmt.Println("A")
	// } else if nilai >= 75 {
	// 	fmt.Println("B")
	// } else if nilai >= 60 {
	// 	fmt.Println("C")
	// } else {
	// 	fmt.Println("D")
	// }

	// // switch
	// switch {
	// case nilai >= 90:
	// 	fmt.Println("A")
	// case nilai >= 75:
	// 	fmt.Println("B")
	// case nilai >= 60:
	// 	fmt.Println("C")
	// default:
	// 	fmt.Println("D")
	// }

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
	// var numbers []string
	// numbers = append(numbers, "satu", "dua", "tiga")
	// fmt.Println(numbers[0:2])

	// person := make(map[string]string)
	// person["name"] = "ajitama"
	// person["job"] = "developer"
	// fmt.Println(person)

	// students := map[string]int{
	// 	"Nozami": 90,
	// 	"Fardan": 70,
	// }

	// fmt.Println(students)
}
