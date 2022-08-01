package main

import (
	"fmt"
	"learn/Book_pro_go/composition/store"
)

func main()  {

	// kayak := store.NewProduct("Kayak", "Watersport", 275)
	// lifejacket := &store.Product{Name: "Lifejacket", Category: "Watersport"}

	// for _, p :=range []*store.Product{kayak, lifejacket}{
	// 	fmt.Println(p.Name, p.Category, p.Price(0.2))
	// }

	// boats := []*store.Boat{
	// 	store.NewBoat("Kayak", 275, 1, false),
	// 	store.NewBoat("Canoe", 400, 3, false),
	// 	store.NewBoat("Tender", 650.25, 2, true),
	// }

	// for _, b := range boats{
	// 	fmt.Println(b.Product.Name, b.Name, b.Price(0.2))
	// }

	// rentals := []*store.RentalBoat{
	// 	store.NewRentalBoat("Rubber ring", 10, 1, false, false, "N/A", "N/A"),
	// 	store.NewRentalBoat("Yacht", 50000, 5, true, true, "Bob", "ALice"),
	// 	store.NewRentalBoat("Super Yacht", 1000000, 15, true, true, "Dora", "Charlie"),
	// }

	// for _, r := range rentals{
	// 	fmt.Println(r.Name, r.Price(0.2), r.Captain)
	// }

	// product := store.NewProduct("Kayak", "Ws", 279)

	// deal := store.NewSpecialDeal("Wee", product, 50)

	// Name, price, Price := deal.GetDetails()
	
	// fmt.Println(Name, price, Price)

	products := map[string]store.ItemForSale{
		"Kayak": store.NewBoat("Kayak", 279, 1, false),
		"Ball": store.NewProduct("Soccer Ball", "Soccer", 19.50),
	}

	for key, p := range products{
		fmt.Println(key, p.Price(0.2))
	}
}