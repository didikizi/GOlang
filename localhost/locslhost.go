package main

import (
	"flag"
	"fmt"
	"net/http"
)

type Config struct {
	Port  string
	Name  string
	Count int
}

func runServer(config Config) {
	http.HandleFunc("/home", func(writer http.ResponseWriter, request *http.Request) {
		data := "Страница приветствует пользователя имя которого было передано при запуске в том количестве которое было указано\n"
		for index := 0; index < config.Count; index++ {
			data += fmt.Sprintf("Hello, %s ! \n", config.Name)
		}
		_, err := fmt.Fprint(writer, data)
		if err != nil {
			return
		}
	})
	http.HandleFunc("/me", func(writer http.ResponseWriter, request *http.Request) {
		data := "Привет!\nЭта страница была создана в  рамках обучения программированию на языке Golang.\nВсе мои проекты расположены по ссылке https://github.com/didikizi/GOlang"
		_, err := fmt.Fprint(writer, data)
		if err != nil {
			return
		}
	})

	err := http.ListenAndServe("localhost:"+config.Port, nil)
	if err != nil {
		return
	}
}

func main() {
	var port = flag.String("port", "8080", "Port number")
	var name = flag.String("name", "Go student", "Name for hello")
	var count = flag.Int("count", 1, "Number of repeats")
	flag.Parse()

	config := Config{
		Port:  *port,
		Name:  *name,
		Count: *count,
	}
	runServer(config)
}
