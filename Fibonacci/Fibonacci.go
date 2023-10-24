package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	number := Read()
	numberFub := Search(number)
	fmt.Println(numberFub)
}

func Read() float64 {
	for {
		fmt.Println("Введите порядковый номер числа Фибоначчи")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()
		number, err := strconv.ParseFloat(line, 64)
		if err == nil {
			return number
		} else {
			fmt.Println("Было введено не коректное значение")
		}
	}
}

func Search(number float64) float64 {

	if number > 2 {
		return (Search(number-1) + Search(number-2))
	} else {
		return 1
	}
}
