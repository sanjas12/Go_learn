package main

import "fmt"

func main() {
	var a, b, sum int
	for fmt.Scan(&a); a == 0; a-- {
		if fmt.Scan(&b); 9 < b && b < 100 && b%8 == 0 {
			
			sum += b
		}
	}
	fmt.Println(sum)
}
