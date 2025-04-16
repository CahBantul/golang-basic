package main

import "fmt"

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
func main() {
	x := 10
	y := 20

	fmt.Println("sebelum", x, y)

	swap(&x, &y)
	fmt.Println("sesudah", x, y)

}
