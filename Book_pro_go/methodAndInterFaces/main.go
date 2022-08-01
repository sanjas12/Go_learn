package main

import "fmt"

type Expence interface {
	getName() string
	getCost(annual bool) float64
}


func calcTotal(expences []Expence) (total float64) {
	for _, item := range expences {
		total += item.getCost(true)
	}
	return total
}

type Account struct {
	accountNumber int
	expenses []Expence
}

func main()  {
	
	account := Account {
		accountNumber: 12345,
		expenses: []Expence{
			&Product{"Kayak","Watersports",275},
			Service{"Boat", 12, 89.5},
		},
	}



	for _, expence := range account.expenses {
		fmt.Println(expence.getName(), expence.getCost(true))
	}

	fmt.Println("Total", calcTotal(account.expenses))

}