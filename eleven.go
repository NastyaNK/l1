package main

import "fmt"

func main() {
	m := []int{2, 3, 4, 5, 6, 9}
	n := []int{3, 1, 7, 9, 5, 4}
	mn := []int{}
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(n); j++ {
			if m[i] == n[j] { //если число из одного массива есть в другом
				mn = append(mn, m[i]) //добавляем в новый
			}
		}
	}
	fmt.Println(mn)
}
