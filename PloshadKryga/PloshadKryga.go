package main

import (
	"fmt"
	"math"
)

func main() {
	var ploch float64
	var radius float64
	fmt.Println("Ввдети площадь круга")
	fmt.Scanln(&ploch)
	radius = math.Sqrt(ploch / math.Pi)
	diametr := radius * 2
	fmt.Println("Диаметр окружности", diametr)
	dlina := diametr * math.Pi
	fmt.Println("Длина окружности", dlina)
}
