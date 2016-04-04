package main

import (
	"fmt"
)

func main() {
	loops := 1
	for loops > 0 {
		fmt.Printf("\nNumber of loops?\n")
		fmt.Scanf("%d", &loops)
		for i := 0; i < loops; i++ {
			fmt.Printf("%d", i)
		}
	}
}
