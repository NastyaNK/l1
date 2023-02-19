package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	record := make(chan int, 10)
	var n int // время
	fmt.Scan(&n)
	go func() {
		for {
			record <- rand.Int() //отправление данных в канал
		}
	}()
	go func() {
		for {
			fmt.Println(<-record) //чтение из канала
		}
	}()
	time.Sleep(time.Duration(n) * time.Second) //остановка программы по истечение n секунд
}
