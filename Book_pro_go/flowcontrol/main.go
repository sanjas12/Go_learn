package main

import ("fmt"
		// "strconv"
)

func main()  {
	// kayak := "275"
	// if kayakPrice, err :=strconv.Atoi(kayak); err == nil {
	// 	fmt.Println(kayakPrice)
	// } else {
	// 	fmt.Println(err)
	// }

		
		// for counter:=0; counter <= 7; counter++ {
		// 	if (counter == 1){
		// 		continue
		// 	}
		// 	fmt.Println(counter)
			// counter++
			// if counter > 3 {
			// 	break
			// }
		// }

		// product := "kayak"
		// product := []string {"Kayak", "Lifejacket", "SoccerBall"}

		// for _, item := range product {
		// 	fmt.Println((item))
		// }
	
		// product := "Kayak"
		// for index, character := range product {
		// 	switch (character) {
		// 		case 'K':
		// 			// if character == 'k' {
		// 			// 	fmt.Println("k at position", index)
		// 			// }
		// 			fmt.Println("Upper K at position", index)
		// 			// fallthrough
		// 		case 'k':
		// 			fmt.Println("lower k at position", index)
		// 		case 'y':
		// 			fmt.Println("y at position", index)
		// 		default:
		// 			fmt.Println("Character", string(character),
		// 			"at position", index)
		// 	}
		// }

		// for counter := 0; counter < 20; counter++ {
		// 	switch val :=counter / 2; val {
		// 		case 2, 3, 5, 7:
		// 			fmt.Println("Prime value:", val)
		// 		default:
		// 			fmt.Println("Non-prime value:", val)
		// 		}
		// }

		// for counter :=0; counter<10; counter++{
		// 	switch {
		// 		case counter == 0:
		// 			fmt.Println("0")
		// 		case counter < 3:
		// 			fmt.Println("меньше 3")
		// 		case counter >=3 && counter < 7:
		// 			fmt.Println("....")
		// 		default:
		// 			fmt.Println("Все остальное")
				
		// 	}
		// }

			counter := 0
			target: fmt.Println("Counter: ", counter)
			counter++
			if (counter < 5){
				goto target
			}


}