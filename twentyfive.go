package main

import (
	"time"
)

func main() {
	sleep(5)
}
func sleep(second float64) {
	data := time.Now()                            // запоминаем текущее время
	for time.Now().Sub(data).Seconds() < second { //отнимает из времени текущее, пока между ними не будет указанная разница в секундах
	}

}
