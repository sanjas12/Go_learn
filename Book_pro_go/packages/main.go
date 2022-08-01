package main

import (
	"fmt"
	// . "learn/Book_pro_go/packages/fmt"
	_ "learn/Book_pro_go/packages/data"
	Fmt "learn/Book_pro_go/packages/fmt"
	"learn/Book_pro_go/packages/store"
	"learn/Book_pro_go/packages/store/cart"

	Color "github.com/fatih/color"
)

func main()  {
	product := store.NewProduct("Kayak", "Watersport", 279)
	fmt.Println("Name:", product.Name)
	fmt.Println("Category:", product.Category)
	// fmt.Println("Price:", product.GetPrice())
	// fmt.Println("Price:", ToCurrency(product.GetPrice()))
	fmt.Println("Price:", Fmt.ToCurrency(product.GetPrice()))


	cart :=cart.Cart {
		CustomerName: "Alice",
		Products: []store.Product {*product},
	}

	fmt.Println("Name cart:", cart.CustomerName)
	fmt.Println("Total:", Fmt.ToCurrency(cart.GetTotal()))

	Color.Green("Name" + cart.CustomerName)
	Color.Cyan("Total" + Fmt.ToCurrency(cart.GetTotal()))

}