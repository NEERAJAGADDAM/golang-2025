package inventory

import (
	"Inventory-Management/product"
	"fmt"
)

func ListProducts() { 
	if len(product.Products) == 0 { 
		fmt.Println("No products available.")
		return
	}

	fmt.Println("\nList of Products:")
	for _, p := range product.Products { 
		fmt.Printf("ID: %d, Name: %s, Category: %s, Price: %.2f, Quantity: %d\n",
			p.ID, p.Name, p.Category, p.Price, p.Quantity)
	}
}
