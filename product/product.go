package product

import (
	//"bufio"
	"fmt"
	//"os"
	//"strings"
)

type Product struct {
	ID       int
	Name     string
	Category string
	Price    float64
	Quantity int
}

var Products = make(map[int]Product)

func AddProduct() {
	var id int
	var name, category string
	var price float64
	var quantity int

	fmt.Println("Enter Product ID:")
	fmt.Scanln(&id)
	fmt.Println("Enter Product Name:")
	fmt.Scanln(&name)
	fmt.Println("Enter Product Category:")
	fmt.Scanln(&category)
	fmt.Println("Enter Product Price:")
	fmt.Scanln(&price)
	fmt.Println("Enter Product Quantity:")
	fmt.Scanln(&quantity)

	product := Product{
		ID:       id,
		Name:     name,
		Category: category,
		Price:    price,
		Quantity: quantity,
	}

	fmt.Println(product)

	Products[id] = product

	fmt.Println("Product added successfully!")
}

func UpdateProduct(id int, Price float64, quantity int) {

	//fmt.Println("Enter product id to update the product:")
	//fmt.Scanln(&id)
	if product, exists := Products[id]; exists {
		//fmt.Println("Price:")
		//fmt.Scanln(&product.Price)
		//fmt.Println("quantity:")
		//fmt.Scanln(&product.Quantity)
		product.Price = Price
		product.Quantity = quantity
		Products[id] = product
		fmt.Println("Product Updated successfully!")
	} else {
		fmt.Println("Product id is not exists in Products inventory")
	}
}

func DeleteProduct(id int) {

	//fmt.Println("Enter product id to remove the product from products:")
	//fmt.Scanln(&id)

	if _, exists := Products[id]; exists {
		delete(Products, id)
		fmt.Println("Product Delete successfully!")
	} else {
		fmt.Println("Product is not found")
	}
}

func SearchProduct() {
	var searchQuery string

	fmt.Println("Enter product name or category to search:")
	fmt.Scanln(&searchQuery)

	found := false
	fmt.Println("\nSearch Results:")
	for _, product := range Products {
		if product.Name == searchQuery || product.Category == searchQuery {
			fmt.Printf("ID: %d, Name: %s, Category: %s, Price: %.2f, Quantity: %d\n",
				product.ID, product.Name, product.Category, product.Price, product.Quantity)
			found = true
		}
	}

	if !found {
		fmt.Println("No matching products found.")
	}
}
