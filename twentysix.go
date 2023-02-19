package main

import (
	"fmt"
	"strings"
)

func s(string2 string) bool {
	nk := map[rune]bool{}
	n := []rune(strings.ToLower(string2)) //переводим все буквы строки в нижний регистр и преобразуем её в массив рун

	for i := 0; i < len(n); i++ {
		_, ok := nk[n[i]]
		if ok { //если в мапе уже была такая руна, то возвращаем false
			return false
		}
		nk[n[i]] = true //если нет добавляем в мппу
	}
	return true
}
func main() {
	fmt.Println(s("Слово")) //тест
}
