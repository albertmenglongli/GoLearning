package main

import (
	"fmt"
)

const (
	Red             = (1 << iota)
	Green           = (1 << iota)
	Blue, ColorMask = (1 << iota), (1 << (iota + 1)) - 1
)
const (
	i complex128 = complex(0, 1)
)

func main() {
	fmt.Printf("Red=", Red)             //1
	fmt.Printf("Green=", Green)         //2
	fmt.Printf("Blue=", Blue)           //4
	fmt.Printf("ColorMask=", ColorMask) //7
	fmt.Printf("i= ", i)
}
