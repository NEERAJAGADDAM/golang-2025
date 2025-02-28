package main

import (
	"Inventory-Management/inventory"
	"Inventory-Management/product"
	"fmt"
)

func main() {
	var option int

	for {
		fmt.Println("----- Welcome to Inventory-Management -----")
		fmt.Println("Enter 1 to Add Products.")
		fmt.Println("Enter 2 to Update Products.")
		fmt.Println("Enter 3 to Delete Products.")
		fmt.Println("Enter 4 to Search Products.")
		fmt.Println("Enter 5 to Display the Existing Products.")
		fmt.Println("Enter 6 to Quit.")
		fmt.Print("Enter your Option: ")
		fmt.Scanln(&option)

		var id int
		var price float64
		var quantity int

		switch option {
		case 1:
			product.AddProduct()
		case 2:
			fmt.Print("Enter Product ID to update: ")
			fmt.Scanln(&id)
			fmt.Print("Enter New Price: ")
			fmt.Scanln(&price)
			fmt.Print("Enter New Quantity: ")
			fmt.Scanln(&quantity)
			product.UpdateProduct(id, price, quantity)
		case 3:
			fmt.Print("Enter Product ID to delete: ")
			fmt.Scanln(&id)
			product.DeleteProduct(id)
		case 4:
			product.SearchProduct()
		case 5:
			inventory.ListProducts()
		case 6:
			fmt.Println("Thank you for visiting Inventory-Management!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
