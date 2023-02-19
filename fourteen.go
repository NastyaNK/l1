package main

import "fmt"

func typeof(nk interface{}) {
	switch nk.(type) {
	case int: //определение типа int
		fmt.Println("int")
	case chan int: //определение типа chan
		fmt.Println("chan")
	case bool: //определение типа bool
		fmt.Println("bool")
	case string: //определение типа string
		fmt.Println("string")
	}
}
func main() {
	typeof(6)
	typeof("hello")
}
