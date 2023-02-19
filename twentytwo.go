package main

import (
	"fmt"
	"math"
)

type kl struct {
	num float64
}

func (n *kl) ml(a, b float64) float64 {
	if n.num < a && n.num < b {
		return a * b
	} //умножение
	return 0
}

func (n *kl) del(a, b float64) float64 {
	if n.num < a && n.num < b {
		return a / b
	} //деление
	return 0
}
func (n *kl) pl(a, b float64) float64 {
	if n.num < a && n.num < b {
		return a + b
	} //сложение
	return 0
}
func (n *kl) mn(a, b float64) float64 {
	if n.num < a && n.num < b {
		return a - b
	} //вычитание
	return 0
}
func main() {
	st := kl{math.Pow(2, 20)} //условие при котором должны происходить мат.операции
	d := st.del(111111111111, 8888778776)
	p := st.pl(4, 6)
	m := st.mn(4, 6)
	ml := st.ml(4, 6)
	fmt.Println(ml, d, m, p)
}
