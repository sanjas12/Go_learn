package main

import (
	"fmt"
	// "strconv"
)


func main()  {

	var price  = "€48.95"
	
	// var currency string = string(price[0])
	// // var currencyByte byte = (price[0])
	// var amoungString string = string(price[1:])

	// amount, parErr := strconv.ParseFloat(amoungString, 64)

	// fmt.Println("currency", currency)
	// fmt.Println("len", len(price), price)
	// if parErr == nil {
	// 	fmt.Println(amount)
	// } else {
	// 	fmt.Println("erorr - ", parErr)
	// }

	for index, char := range []byte(price){
		fmt.Println(index, char, string(char))
	}
} 