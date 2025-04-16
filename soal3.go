package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

func UpdatePrice(product *Product) {
	product.Price = 50
}

func main() {
	product := Product{Name: "mouse", Price: 10}
	fmt.Println("harga sebelum", product.Price)

	UpdatePrice(&product)
	fmt.Println("harga sesudah", product.Price)

}
