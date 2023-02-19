package main

import "fmt"

func main() {
	num1 := 1
	num := 2
	num1, num = num, num1 //замена чисел без создания третьей переменнной
	fmt.Println(num1, num)
}
