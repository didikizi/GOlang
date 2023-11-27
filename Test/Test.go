package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"log"
)

type Config struct {
	Port  string
	Name  string
	Count int
}

func runServer(config Config) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		data:= "\n"
		for index := 0; index < config.Count; index++ {
			data += fmt.Sprintf("Hello, %s ! \n", config.Name)
		}
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
	var name = flag.String("NameFail", "", "Port number")
	flag.Parse()
	if name== nil{
		log.Fatalln("No file name")
	}
	f, err := os.OpenFile (*name, os.O_RDONLY,0)
	defer func (){
		err = f.Close()
		if err!= nil{
			log.Fatalln(err)
		}
	}()
	fmt.Println(os.Stat(name))

	var config Config
	var data []byte
	_, err = f.Read(data)
	if err!= nil {
		log.Fatalln(err)
	}
	fmt.Println(data)
	if err!= nil{
		log.Fatalln(err)
	}

	runServer(config)
}
