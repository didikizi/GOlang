package main

import (
	"fmt"
)

func main() {
	var length, width uint8
	fmt.Println("Ввдети длину и ширину")
	fmt.Scanln(&length, &width)
	square := length * width
	fmt.Println(square)
}
