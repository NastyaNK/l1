package main

import (
	"fmt"
	"strings"
)

func main() {
	r := strings.Split("snow dog sun", " ") //делим строку по пробелам
	var s string
	for i := 0; i < len(r); i++ {
		s += r[len(r)-i-1] + " " //добавляем в новые строки из массива в обратном порядке
	}
	fmt.Println(s)

}
