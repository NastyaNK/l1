package main

import (
	"fmt"
	"math/rand"
)

func main() {
	c := make(chan int)
	n := 6 //воркеры
	for i := 0; i < n; i++ {
		go func() {
			for {
				fmt.Println(<-c) //чтение данных из канала с выводом в stdou
			}
		}()
	}
	for { //бесконечный цикл в главном потоке
		c <- rand.Int() //для постоянной записи произвольный чисел в канал
	}
} //все воркеры завершатся вместе с основным потоком
