package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	Mapa := map[int]int{}
	for {
		number := Read()
		numberFub, err := Mapa[number]
		if !err {
			numberFub = Search(number)
			Mapa[number] = numberFub
			fmt.Println(numberFub)
		} else {
			fmt.Println(numberFub)
		}
	}
}

func Read() int {
	for {
		fmt.Println("Введите порядковый номер числа Фибоначчи")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()
		number, err := strconv.Atoi(line)
		if err == nil {
			return number
		} else {
			fmt.Println("Было введено не коректное значение")
		}
	}
}

func Search(number int) int {

	if number > 2 {
		return (Search(number-1) + Search(number-2))
	} else {
		return 1
	}
}
