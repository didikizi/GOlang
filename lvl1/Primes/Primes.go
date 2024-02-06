package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	for {
		fmt.Println("Введите число до которого необходим поиск")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()
		number, err := strconv.Atoi(line)

		if err == nil {
			for index := 1; index != number; index++ {
				if index%2 != 0 && index%3 != 0 && index%5 != 0 && index%7 != 0 {
					fmt.Println(index)
				}
			}
		} else {
			fmt.Println("Введено не корректное число")
		}
	}
}
