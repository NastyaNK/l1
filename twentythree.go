package main

import "fmt"

func ms(sl []int, i int) []int {
	return append(sl[0:i], sl[i+1:]...) //Удаление i элемента из слайса
}
func main() {
	sl := []int{1, 6, 9, 4, 7, 1, 0}
	fmt.Println(ms(sl, 4))
}
