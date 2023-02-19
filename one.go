package main

import "fmt"

type Human struct {
	name   string
	female string
	year   int
}

type Action struct {
	height int
	weight int
	Human  //Аналог наследования
}

func main() {

	result := Action{

		175, 62, Human{"Лена", "Женский", 22}, //Проверка
	}

	fmt.Println(result)
}
