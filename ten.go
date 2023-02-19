package main

import "fmt"

func main() {
	array := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	m := make(map[int][]float32)
	for i := 0; i < len(array); i++ {
		el := int(array[i]/10) * 10 //округление в меньшую сторону
		_, ok := m[el]
		if !ok {
			m[el] = []float32{array[i]} //создание в map массива
		} else {
			m[el] = append(m[el], array[i]) //добавление в map
		}
	}
	fmt.Println(m)
}
