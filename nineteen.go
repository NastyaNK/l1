package main

import "fmt"

func main() {
	r := []rune("главрыба") //переводим в руны для замены символов местами
	for i := 0; i < len(r)/2; i++ {
		r[i], r[len(r)-i-1] = r[len(r)-i-1], r[i] //обратный порядок строки
	}
	fmt.Println(string(r))
}
