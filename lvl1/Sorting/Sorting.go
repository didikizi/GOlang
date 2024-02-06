package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	for {
		flag, matrix := Read()
		if flag {
			fmt.Println("Не отсортированая строка:")
			fmt.Println(matrix)
			fmt.Println("Отсортированая строка:")
			fmt.Println(Sort(matrix))
		} else {
			fmt.Println("Были введены не корректные значения")
		}
	}
}

func Read() (bool, []int) {
	fmt.Println("Введите строку чисел разбеляя их пробелами")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	flag := true
	matrix := []int{}
	words := strings.Fields(line)
	for _, word := range words {
		number, err := strconv.Atoi(word)
		if err != nil {
			flag = false
			return flag, matrix
		} else {
			matrix = append(matrix, number)
		}
	}
	return flag, matrix
}

func Sort(matrix []int) []int {
	for index := 0; index < len(matrix)-1; index++ {
		for index1 := index; index1 < len(matrix); index1++ {
			if matrix[index] < matrix[index1] {
				TMP := matrix[index]
				matrix[index] = matrix[index1]
				matrix[index1] = TMP
			}
		}
	}
	return matrix
}
