package main

import "fmt"

func ubahArray(arr *[5]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = 1
	}
}

func main() {
	data := [5]int{3, 4, 5, 8, 7}

	fmt.Println("sebelum", data)

	ubahArray(&data)
	fmt.Println("sesudah", data)

}
