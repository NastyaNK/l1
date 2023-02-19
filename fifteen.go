package main

import (
	"fmt"
	"strings"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string([]rune(v)[:100]) //использование массива рун
}

func createHugeString(n int) string {
	var sb strings.Builder
	sb.Grow(n) //выделение нужной памяти
	for i := 0; i < n; i++ {
		sb.WriteByte(52) //запись случайного байта указанное количество раз
	}
	return sb.String() //сбор строки
}

func main() {
	someFunc()
	fmt.Println(justString)
}
