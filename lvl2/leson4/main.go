package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	for {
		fmt.Println("1. Счетчик до 10\n2.Запуск всех 10 потоков и завершение их всех")
		var a int
		fmt.Fscan(os.Stdin, &a)
		fmt.Println("Начало")
		switch a {
		case 1:
			case1()
		case 2:
			case2()
		}
	}
}

// Структура Chitдля реализации счетчика в case 1
type Chit struct {
	mu  sync.Mutex
	wg  sync.WaitGroup
	col int
}

// Пустая структура для каналов
type TMP struct{}

// Функция реализует рост счетчика Chit col на единицу каждую секунду и использование горутин формы goworkera и контекста задачи
func (one *Chit) Add(ctxWorker context.Context, mutex chan struct{}, done chan struct{}, index int) {
	for {
		select {
		case <-ctxWorker.Done():
			<-mutex
			return
		case <-mutex:
			if one.col >= 10 {
				done <- TMP{}
				return
			}
			one.col++
			fmt.Println(one.col, "\t Я горутина ", index)
			time.Sleep(1 * time.Second)
			mutex <- TMP{}
		default:
		}
	}
}

// Функция для контроля выполнения задачи и завершения контекста задачи в случае принудительного завершения выполнения
func Worker(sigChan chan os.Signal, done chan struct{}, cancelFunc context.CancelFunc) {
	for {
		select {
		case <-sigChan:
			cancelFunc()
			time.Sleep(1 * time.Second)
			fmt.Println("\nCмерть")
			return
		case <-done:
			cancelFunc()
			fmt.Println("Задача выполнена")
			return
		}
	}
}

// Выполнение счетчика от 1 до 10 через горутины, mutex каналы и контекст задачи
func case1() {
	mutex := make(chan struct{}, 1)    //Канал реализации mutex
	done := make(chan struct{}, 1)     //Канал реализации выполнения задачи
	sigChan := make(chan os.Signal, 1) //Канал для перехвата завершения программы
	signal.Notify(sigChan, os.Interrupt, syscall.SIGINT)
	ctx := context.Background() //Прерываемый контекст задачи
	ctxWorker, cancelFunc := context.WithCancel(ctx)
	one := Chit{col: 0}                   //Структура для счетчика
	mutex <- TMP{}                        //Запуск канала mutex
	for index := 0; index < 10; index++ { //Создания горутин для выполнения задачи
		go one.Add(ctxWorker, mutex, done, index)
	}
	Worker(sigChan, done, cancelFunc) //Модуль отслеживания работы горутин и завершения контекста

	defer close(done)
	defer close(mutex)
	defer signal.Stop(sigChan)
	defer close(sigChan)
}

// Выполнение счетчика от 1 до 10 через горутины и стандартный mutex
func case2() {
	quantity := 10
	two := Chit{col: 0}
	for index := 0; index < quantity; index++ {
		two.wg.Add(1)
		go func(two *Chit, index int) {
			two.mu.Lock()
			defer two.mu.Unlock()
			two.col = two.col + 1
			fmt.Println(two.col, "\tЯ горутина ", index+1)
			two.wg.Done()
		}(&two, index)
	}
	two.wg.Wait()
}
