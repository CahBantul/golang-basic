package main

import "fmt"

func ubahNilai(val *int) {
	*val = 100
}
func main() {
	nilai := 5
	ubahNilai(&nilai)
	fmt.Println("Nilai setelah dirubah:", *&nilai)

}
