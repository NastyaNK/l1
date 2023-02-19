package main

import (
	"fmt"
	"math"
)

type Point struct { //скрытые поля (для других пакетов)
	x int
	y int
}

func main() {
	p := New(6, 4)
	p1 := New(8, 2)
	fmt.Println(p.Distance(&p1))
	fmt.Println(p, p1)
}
func (p *Point) SetY(x int) { //сеттеры для изменения координат
	p.x = x
}
func (p *Point) SetX(y int) {
	p.y = y
}
func (p *Point) GetY() int { // геттеры для получения координат
	return p.y
}
func (p *Point) GetX() int {
	return p.x
}
func New(x, y int) Point {

	return Point{x: x, y: y}
}
func (p *Point) Distance(point *Point) float64 { //функция для расчёта расстояния между двумя точками
	return math.Sqrt(float64(((p.x - point.x) * (p.x - point.x)) + ((p.y - point.y) * (p.x - point.y))))
	return 0
}
