package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup //WaitGroup для параллельной работы
	var Noname int
	for i := 2; i < 10; i += 2 { //Проход по чётным числам до 10
		wg.Add(1)
		go func(i int) { //запуск потока для конкурентного расчёта
			Noname += i * i // сумма квадратов
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(Noname)
}
