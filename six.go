package main

import "time"

func s1() { //1
	ch := make(chan int)
	go func() {
		for {
			data, ok := <-ch
			if !ok {
				return
			}
			println(data)
		}
	}()
	ch <- 123
	close(ch)
}

func main() { //2
	ch := make(chan int)
	go func() {
		for {
			select {
			case <-ch:
				return
			default:
				println("work...")
			}
		}
	}()
	time.Sleep(1 * time.Second)
	close(ch)
}
