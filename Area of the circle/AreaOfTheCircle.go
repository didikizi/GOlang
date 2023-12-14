package main

import (
	"fmt"
	"math"
)

func main() {
	var square float64
	var radius float64
	fmt.Println("Ввдети площадь круга")
	fmt.Scanln(&square)
	radius = math.Sqrt(square / math.Pi)
	Diameter := radius * 2
	fmt.Println("Диаметр окружности", Diameter)
	Length := Diameter * math.Pi
	fmt.Println("Длина окружности", Length)
}
