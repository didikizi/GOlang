package main

import (
	"fmt"
)

func main() {
	var numders string
	fmt.Println("Введи трехзначное число") // Более крутой способ
	fmt.Scanln(&numders)
	fmt.Println("Единицы", string(numders[2]))
	fmt.Println("Десятки", string(numders[1]))
	fmt.Println("Сотни", string(numders[0]))

	var numbers2 uint16
	fmt.Println("Введи трехзначное число") // Оптимальное решение
	fmt.Scanln(&numbers2)
	hundred := numbers2 / 100
	fmt.Println("Сотни", hundred)
	dozen := (numbers2 - (hundred * 100)) / 10
	fmt.Println("Десятки", dozen)
	unit := (numbers2 - (hundred * 100) - (dozen * 10))
	fmt.Println("Единицы", unit)
}
