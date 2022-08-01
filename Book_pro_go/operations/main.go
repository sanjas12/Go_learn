package main

import (
	"fmt"
	"strconv"
	// "strconv"
	// "math"
)

func main()  {
	// price, tax := 275.0, 27.40
	// sum := price + tax
	// difference := price - tax
	// product := price * tax
	// quotient := price / tax
	// fmt.Println(sum)
	// fmt.Println(difference)
	// fmt.Println(product)
	// fmt.Println(quotient)

	// first := &sum
	// second := &sum

	// fmt.Println(*first==*second)

// 	kayak := 275
// soccerBall := 19.50
// total := float64(kayak) + soccerBall
// fmt.Println(total)


// kayak := 275
// soccerBall := 19.50
// total := kayak + int(soccerBall)
// fmt.Println(total)
// fmt.Println(int8(total))
// fmt.Println(kayak + int(math.Round(10.49)))

// val1 := "true"
// val2 := "false"
// val3 := "not true"
// bool1, b1err := strconv.ParseBool(val1)
// bool2, b2err := strconv.ParseBool(val2)
// bool3, b3err := strconv.ParseBool(val3)
// fmt.Println("Bool 1", bool1, b1err)
// fmt.Println("Bool 2", bool2, b2err)
// fmt.Println("Bool 3", bool3, b3err)


// val1 := "01"
	// val1 := "-1000"

	// if bool1, b1err :=strconv.ParseInt(val1, 10, 64); b1err == nil {
	// 	fmt.Println("Зчанение переменной: ", bool1)
	// } else {
	// 	fmt.Println("No correct", val1 )		
	// }

	// val1 :="10000000"
	// val1 :="10023000000000"

	// // int1, err1 :=strconv.ParseInt(val1, 0, 0)
	// int1, err1 :=strconv.Atoi(val1)

	// if err1 == nil {
	// 	fmt.Println("Value is: ", int1)
	// } else {
	// 	fmt.Println("Problem", val1, err1)
	// }

	// val1 := "4.895e+01"
	// float1, float1err := strconv.ParseFloat(val1, 64)
	// if float1err == nil {
	// 	fmt.Println("Parsed value:", float1)
	// } else {
	// 	fmt.Println("Cannot parse", val1, float1err)
	// }

	var1 := 23.4232

	float1 :=strconv.FormatFloat(var1, 'f', 0, 32)
		fmt.Println(float1)
	}		