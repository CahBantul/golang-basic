package main

import "fmt"

func main() {
	nilai := 5
	var ptr *int = &nilai
	fmt.Println("Alamat dari angka:", ptr)
	fmt.Println("Nilai dari pointer:", *ptr)

}
