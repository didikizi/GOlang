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
		fmt.Println(Mapa)
		numberFub := Search(number, Mapa)
		fmt.Println(numberFub)
		fmt.Println(Mapa)
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

func Search(number int, Mapa map[int]int) int {

	if number > 2 {
		_, err := Mapa[number]
		if err {
			return Mapa[number]
		} else {
			Mapa[number] = (Search(number-1, Mapa) + Search(number-2, Mapa))
			return Mapa[number]
		}
	} else {
		return 1
	}
}
