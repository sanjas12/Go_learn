package cart

import (
	"learn/Book_pro_go/packages/store"
)

type Cart struct {
	CustomerName string
	Products []store.Product 
}

func (cart *Cart) GetTotal() (total float64) {
	for _, p := range cart.Products {
		total += p.GetPrice()
	}
	return
}