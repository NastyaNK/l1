package main

import (
	"fmt"
	"math/rand"
)

func quicksort(ms []int) []int {
	if len(ms) < 2 { // Если массив состоит из одного элемента или вообще пустой заканчиваем сортировку
		return ms
	}
	left, right := 0, len(ms)-1                 // Определяем границы массива
	pivot := rand.Int() % len(ms)               // Выбираем случайным образом опорную точку
	ms[right], ms[pivot] = ms[pivot], ms[right] // Меняем её с самым правым элементом
	for i := range ms {                         // Проходим по массиву
		if ms[i] < ms[right] { // Все элементы которые меньше опорной точки
			ms[left], ms[i] = ms[i], ms[left] // Переносим левее
			left++
		}
	}
	ms[left], ms[right] = ms[right], ms[left] // Возвращаем опорнуб точку на свое место
	quicksort(ms[:left])                      // Запускаем сортировку массива левее от опорной точки
	quicksort(ms[left+1:])                    // Запускаем сортировку правее
	return ms
}

func main() {
	ms := []int{4, 3, 124, 33, 213, 4, 3, 5, 8, 0, 1, 3, 35, 78}
	fmt.Println(quicksort(ms))
}
