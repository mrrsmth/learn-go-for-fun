package types

import (
	"fmt"

)

type mile int
type kilometer int
type library []string

// type person struct {
// 	name string
// 	age  int
// }

type contact struct {
	email string
	phone string
}

type person struct {
	name string
	age  int
	contact
}

func (p person) print() {
	fmt.Println("Имя:", p.name)
	fmt.Println("Возраст:", p.age)
}

func (p person) eat(meal string) {
	fmt.Println(p.name, "ест", meal)
}

type node struct {
	value int
	next  *node
}

func printNodeValue(n *node) {

	fmt.Println(n.value)
	if n.next != nil {
		printNodeValue(n.next)
	}
}

func printBooks(lib library) {

	for _, value := range lib {

		fmt.Println(value)
	}
}

func distanceToEnemy(distance mile) {

	fmt.Println("расстояние для противника:")
	fmt.Println(distance, "миль")
}

func (p *person) updateAge(newAge int) {
	(*p).age = newAge
}

func TypesMain() {
	var tom = person{name: "Tom", age: 24}
	var tomPointer *person = &tom
	fmt.Println("before", tom.age)
	tomPointer.updateAge(33)
	fmt.Println("after", tom.age)

	// 1
	// var distance mile = 5
	// distance := mile(5)
	// distanceToEnemy(distance)

	// var myLibrary library = library{"Book1", "Book2", "Book3"}
	// printBooks(myLibrary)

	// var tom = person{name: "Tom", age: 24}
	// fmt.Println(tom.name) // Tom
	// fmt.Println(tom.age)  // 24

	// tom.age = 38 // изменяем значение
	// fmt.Println(tom.name, tom.age)
	// 2
	// tom := person{name: "Tom", age: 22}
	// // var tomPointer *person = &tom
	// tomPointer := &tom
	// valueType := reflect.TypeOf(tomPointer)
	// valueTypeSecond := reflect.TypeOf(tom)
	// // tomPointer.age = 29
	// // fmt.Println(tom.age) // 29
	// // (*tomPointer).age = 32
	// // fmt.Println(tom.age)
	// // value := 42
	// // valueType := reflect.TypeOf(value)
	// fmt.Println(valueType)
	// fmt.Println(valueTypeSecond)

	// var tom = person{
	// 	name: "Tom",
	// 	age:  24,
	// 	contact: contact{
	// 		email: "tom@gmail.com",
	// 		phone: "+1234567899",
	// 	},
	// }
	// tom.contact.email = "supertom@gmail.com"

	// fmt.Println(tom.contact.email) // supertom@gmail.com
	// fmt.Println(tom.contact.phone)
	// tom.print()
	// tom.eat("борщ с капустой, но не красный")

	// first := node{value: 4}
	// second := node{value: 5}
	// third := node{value: 6}

	// first.next = &second
	// second.next = &third

	// var current *node = &first
	// for current != nil {
	// 	fmt.Println(current.value)
	// 	current = current.next
	// }
}
