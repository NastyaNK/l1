package main

import (
	"fmt"
	"sync"
)

type st struct {
	count int
	mu    sync.Mutex
}

func main() {
	var wg sync.WaitGroup //для синхронизации потоков
	var n st
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			n.mu.Lock()   //блокирует чтобы избежать гонки данных
			n.count++     //счетчик
			n.mu.Unlock() //убираем блокировку
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(n.count)
}
