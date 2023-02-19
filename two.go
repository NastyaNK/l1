package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup           //WaitGroup для параллельной работы
	arr := []int{2, 4, 6, 8, 10}    //числа  которые возведем в квадрат
	for i := 0; i < len(arr); i++ { //проход по массиву
		wg.Add(1)
		go func(i int) { //запуск потока для конкурентного расчёта
			fmt.Println(i * i) //вывод квадрата в stdout
			wg.Done()
		}(arr[i])
	}
	wg.Wait()
}
