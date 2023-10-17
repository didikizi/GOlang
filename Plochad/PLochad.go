package main

import (
	"fmt"
)

func main() {
	var dlina, shirina uint8
	fmt.Println("Ввдети длину и ширину")
	fmt.Scanln(&dlina, &shirina)
	ploshad := dlina * shirina
	fmt.Println(ploshad)
}
