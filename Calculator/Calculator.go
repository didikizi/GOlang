package main

import (
	"fmt"
	//"errors"
	//"math"
)

// Работает не коректно из за того что любой err который произошел более одного раза вызывает создание псевдоцикла
func main() {
	var x, y int8
	var operant string
	flag := true
	for {
		fmt.Println("Введите X:")
		for {
			_, err := fmt.Scan(&x)
			if err != nil {
				fmt.Println("Введено не коректное значение")
			} else {
				err = nil
				break
			}
		}
		fmt.Println("Введите Y:")
		for {
			_, err := fmt.Scan(&y)
			if err != nil {
				fmt.Println("Введено не коректное значение")
			} else {
				break
			}
		}
		fmt.Println("Введите действие:")
		fmt.Scan(&operant)
		for flag {
			switch operant {
			case "+":
				answer := x + y
				fmt.Println("Ответ :", answer)
				flag = false
			case "-":
				answer := x - y
				fmt.Println("Ответ :", answer)
				flag = false
			case "*":
				answer := float32(x) * float32(y)
				fmt.Println("Ответ :", answer)
				flag = false
			case "/":
				answer := float32(x) / float32(y)
				fmt.Println("Ответ :", answer)
				flag = false
			default:
				fmt.Println("Введена некоректная операция, введите заново")
				fmt.Scan(&operant)
			}
		}
	}
}
