package main

import "fmt"

func main() {
	n := []string{"cat", "cat", "dog", "cat", "tree"} //Добавляем элементы из массива в мапу как ключи для избавления от повторов
	s := make(map[string]interface{})
	for _, s2 := range n {
		s[s2] = nil
	}
	for s2, _ := range s {
		fmt.Println(s2)
	}
}
