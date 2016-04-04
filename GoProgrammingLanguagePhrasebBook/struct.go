package main

import (
	"fmt"
)

type Example struct {
	Val   string
	count int
}

func log(i int) {
	fmt.Printf("%d\n", i)
}

func Log(e *Example) {
	e.count++
	fmt.Printf("%d %s\n", e.count, e.Val)
}
func main() {
	log(1)
	// Log()
}
