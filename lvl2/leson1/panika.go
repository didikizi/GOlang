package main

import (
	"fmt"
	"log"
	"time"
)

type Errors struct {
	message string
	errtime time.Time
}

func main() {
	var a int
	fmt.Println("Добрый день!")
	fmt.Scan(&a)
	tmp, err := One(a)
	if err.message != "" {
		log.Fatalln(err.message, err.errtime)
	}
	fmt.Println(tmp)
}

func One(a int) (int, Errors) {
	err := Errors{}
	defer func() {
		if v := recover(); v != nil {
			err = Errors{"Паника", time.Unix(time.Now().Unix(), 0).UTC()}
			fmt.Println(err)
		}
	}()
	if a == 1 {
		err = Errors{"Ошибка а = 1", time.Unix(time.Now().Unix(), 0).UTC()}
		return 0, err
	}
	b := 10
	ab := b / a
	return ab, err
}
