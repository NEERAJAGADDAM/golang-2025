package main

import (
<<<<<<< HEAD
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
=======
	"fmt"
	"LRU-Problem/put"
	"LRU-Problem/get"
)

func main() {
	put.Put(1, 10)
	put.Put(2, 20)
	put.Put(3, 30)
	put.Put(4, 40)

	fmt.Println("Getting Value for key 2:", get.Get(2))
	fmt.Println("Getting Value for key 1:", get.Get(1))
	fmt.Println("Getting Value for key 3:", get.Get(3))
>>>>>>> 7d8a5d5 (Second  commit)
}
