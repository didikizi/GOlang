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
		fmt.Println("1. Счетчик до 10 с проблемами возведенными в абсолют\n2. Счетчик до 10 нормального человека")
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

// Выполнение счетчика от 1 до 10 через горутины без mutex без wg group
// Так же без использования контекста перехвата сигнала
// Минусом такого решения является то что не получиться закрыть за собой каналы
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

	defer signal.Stop(sigChan)
	defer close(sigChan)
}

// Выполнение счетчика от 1 до 10 с использованием всех сложных функций
// Минусов нет, но сложно(
func case2() {
	two := Chit{col: 0}

	ctx := context.Background()
	signalCtx, closed := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGABRT)
	defer closed()

	mutex := make(chan struct{}, 1)
	defer close(mutex)
	mutex <- TMP{}

	for index := 0; index < 10; index++ {
		two.wg.Add(1)
		go two.Add2(mutex, signalCtx)
	}

	two.wg.Wait()
	fmt.Println("Задача выполнена")
}

// For case 1
// Функция для контроля выполнения задачи и завершения контекста задачи в случае принудительного завершения выполнения
func Worker(sigChan chan os.Signal, done chan struct{}, cancelFunc context.CancelFunc) {
	for {
		select {
		case <-sigChan:
			cancelFunc()
			fmt.Println("\nCмерть")
			return
		case <-done:
			cancelFunc()
			fmt.Println("Задача выполнена")
			return
		}
	}
}

// Структура Chit реализации счетчиков
type Chit struct {
	mu  sync.Mutex
	wg  sync.WaitGroup
	col int
}

// Пустая структура для каналов
type TMP struct{}

// For case 1
// Функция реализует рост счетчика Chit col на единицу каждую секунду и использование горутин формы goworkera и контекста задачи
func (one *Chit) Add(ctxWorker context.Context, mutex chan struct{}, done chan struct{}, index int) {
	for {
		select {
		case <-ctxWorker.Done():
			return
		case <-mutex:
			if one.col >= 10 {
				done <- TMP{}
				mutex <- TMP{}
				return
			}
			one.col++
			fmt.Println(one.col, "\t Я горутина ", index+1)
			time.Sleep(1 * time.Second)
			mutex <- TMP{}
		default:
		}
	}
}

// For case 2
// Функция реализует счетчик через констекс сигналов и один канал для сохдания последовательного вывода
func (two *Chit) Add2(mutex chan struct{}, signalCtx context.Context) {
	defer two.wg.Done()
	for {
		select {
		case <-signalCtx.Done():
			return
		case <-mutex:
			if two.col <= 9 {
				two.col++
				fmt.Println(two.col)
				time.Sleep(1 * time.Second)
				mutex <- TMP{}
			} else {
				mutex <- TMP{}
				return
			}
		}
	}
}
