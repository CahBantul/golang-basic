package main

import "fmt"

func main() {
	var ptr *string

	if ptr != nil {
		fmt.Println("Nilai", *ptr)
	} else {
		fmt.Println("pointer masih nil")
	}
}
