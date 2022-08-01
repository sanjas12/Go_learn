package main

import (
	"fmt"
	"os"
)

	func main() {
		// var s, sep string
		
		// for i:=0; i<len(os.Args); i++{
		// 	s +=sep + os.Args[i]
		// 	sep = " "
		// }

			for ind, arg := range os.Args[:]{
				// s +=sep + arg
				// sep = " "
				fmt.Println(ind, arg)
			}
		
		// fmt.Println(s)
		// fmt.Println(strings.Join(os.Args[1:], " "))
		// fmt.Println(os.Args[1:])
		// fmt.Println("Hello мир")
		// fmt

		// fmt.Println(string(time.Now()))
	}
