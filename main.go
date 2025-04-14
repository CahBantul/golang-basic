package main

import "fmt"

// func sayHello(name string) {
// 	fmt.Println("hallo", name)
// }

// func tambah(a, b int) int {
// 	return a + b
// }

//	func calculate(a, b int) (sum int, multiply int) {
//		sum = a + b
//		multiply = a * b
//		return
//	}
// func handlePanic() {
// 	if r := recover(); r != nil {
// 		fmt.Println("panic", r)
// 	}
// }

// func safeDivision(a, b int) {
// 	defer handlePanic()

// 	if b == 0 {
// 		panic("pembagian dengan nol tidak diperbolehkan")
// 	}

// 	result := a / b
// 	fmt.Println("hasil Pembagian", result)
// }

// struct
// type User struct {
// 	Name    string
// 	Email   string
// 	Age     int
// 	Address Address
// }

// type Address struct {
// 	City    string
// 	Country string
// }

// // struct method
// func (user User) IsAdult() bool {
// 	return user.Age > 18
// }

type Speaker interface {
	Speak() string
}

type Dog struct{}

func (dog Dog) Speak() string {
	return "Guk guk!"
}

type Cat struct{}

func (cat Cat) Speak() string {
	return "Meong Meong"
}

func main() {
	dog := Dog{}
	result := dog.Speak()
	fmt.Println("hewan ini bersuara", result)
	cat := Cat{}
	catSpeak := cat.Speak()
	fmt.Println("hewan ini bersuara", catSpeak)
	// user1 := User{
	// 	Name:  "Ajitama",
	// 	Email: "ajitama@gmail.com",
	// 	Age:   33,
	// 	Address: Address{
	// 		City:    "Surabaya",
	// 		Country: "Indonesia",
	// 	},
	// }
	// user2 := User{
	// 	Name:  "nozami",
	// 	Email: "ajitama@gmail.com",
	// 	Age:   13,
	// }

	// fmt.Println(user1)
	// fmt.Println(user1.Name)
	// fmt.Println(user1.IsAdult())
	// fmt.Println(user2.IsAdult())

	// defer handlePanic()
	// fmt.Println("start")
	// safeDivision(1, 0)
	// safeDivision(10, 2)
	// fmt.Println()
	// sayHello("Ajitama")
	// result := tambah(5, 10)
	// fmt.Println("hasil pemjumlahan", result)
	// add, multiply := calculate(3, 6)
	// fmt.Println("hasil pemjumlahan", add)
	// fmt.Println("hasil perkalian", multiply)

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
