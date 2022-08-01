package main

import (
	"fmt"
	"sort"
)


func main() {
	names := [3]string{"fa", "b22222", "c"}

	secondName := &names[1]

	fmt.Println(*secondName, names[1])

	sort.Strings(names[:])
	

	fmt.Println(*secondName, names[1])

	}