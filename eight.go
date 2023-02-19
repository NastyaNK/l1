package main

import (
	"fmt"
	"strconv"
)

func set(x int64, i int, bit int) int64 {
	if bit == 0 {
		return x ^ (1 << (i - 1))
	} //Создаем маску с нужным битом и применяем исключающее ИЛИ (XOR)
	return x | (1 << (i - 1)) //Создаем маску и применяем ИЛИ
}

func main() {
	var x int64
	x = 167
	i := 4
	bit := 1
	y := set(x, i, bit)
	fmt.Printf("%d -> %d\n", x, y)
	fmt.Printf("%s -> %s\n", strconv.FormatInt(x, 2), strconv.FormatInt(y, 2))
}
