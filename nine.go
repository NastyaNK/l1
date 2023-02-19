package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(chan int)
	c := make(chan int)
	array := []int{2, 4, 6, 8}
	var wg sync.WaitGroup
	go func() {
		for {
			c <- <-m * 2 // во второй — результат операции
		}
	}()
	go func() {
		for {
			fmt.Println(<-c) //получаем данные из канала в stdout
			wg.Done()
		}
	}()
	wg.Add(len(array))
	for i := 0; i < len(array); i++ {
		m <- array[i] //записывпаем числа в первый канал
	}
	wg.Wait()
}
