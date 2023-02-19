package main

import "fmt"

type Animal interface { //интерфейс животного который используется в функции
	name() string
	eat(string)
	sound()
}

type Dog struct{} // структура собаки реализующая интерфейс

func (dog *Dog) name() string {
	return "Собака"
}
func (dog *Dog) eat(eda string) {
	fmt.Println(dog.name(), "сегодня ест "+eda)
}
func (dog *Dog) sound() {
	fmt.Println(dog.name(), "говорит гав")
}

type Cat struct { // структура кота реализующая интерфейс
}

func (cat *Cat) eat(eda string) {
	fmt.Println(cat.name(), "сегодня ест "+eda)
}
func (cat *Cat) name() string {
	return "Кот"
}
func (cat *Cat) sound() {
	fmt.Println(cat.name(), "говорит мяу")
}

type RoboDog struct{} //робот пёс который не реализует интерфейс

func (rdog *RoboDog) name() string {
	return "Робот собака"
}
func (rdog *RoboDog) sound() {
	fmt.Println(rdog.name(), "привет")
}
func (rdog *RoboDog) charge() {
	fmt.Println(rdog.name(), "заряжается")
}

type Adapter struct { //адаптер для использования робота пса как интерфейс животного, который его реализует
	robo RoboDog
}

func (adapter *Adapter) name() string {
	return adapter.robo.name()
}
func (adapter *Adapter) eat(eda string) {
	fmt.Println(adapter.name(), "не ест", eda)
	adapter.robo.charge()
}
func (adapter *Adapter) sound() {
	adapter.robo.sound()
}

func walk(animal Animal) { //прогулка животного
	fmt.Println("Сегодня гуляет", animal.name())
	animal.eat("Корм")
	animal.sound()
}

func main() { //тестирование
	dog := Dog{}
	walk(&dog)

	cat := Cat{}
	walk(&cat)

	robo := RoboDog{}
	walk(&Adapter{robo})
}
