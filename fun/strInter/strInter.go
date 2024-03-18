package strinter

import "fmt"

type Reader interface {
	read()
}

type Writer interface {
	write(string)
}

func writeToStream(writer Writer, text string) {
	writer.write(text)
}
func readFromStream(reader Reader) {
	reader.read()
}

type File struct {
	text string
}

func (f *File) read() {
	fmt.Println(f.text)
}
func (f *File) write(message string) {
	f.text = message
	fmt.Println("Запись в файл строки", message)
}

type contact struct {
	email string
	phone string
}

type person struct {
	name        string
	age         int
	contactInfo contact
}

func (p person) print() {
	fmt.Println("Имя:", p.name)
	fmt.Println("Возраст:", p.age)
}

func (p person) eat(meal string) {
	fmt.Println(p.name, "ест", meal)
}

type Vehicle interface {
	move()
}

// структура "Автомобиль"
type Car struct{}

// структура "Самолет"
type Aircraft struct{}

func (c Car) move() {
	fmt.Println("Автомобиль едет")
}
func (a Aircraft) move() {
	fmt.Println("Самолет летит")
}

func Init() {
	var tom = person{
		name: "Tom",
		age:  24,
		contactInfo: contact{
			email: "tom@gmail.com",
			phone: "+1234567899",
		},
	}
	tom.contactInfo.email = "supertom@gmail.com"
	tom.print()
	tom.eat("борщ с капустой, но не красный")
	fmt.Println(tom.contactInfo.email) // supertom@gmail.com
	fmt.Println(tom)                   // +1234567899

	var tesla Vehicle = Car{}
	var boing Vehicle = Aircraft{}
	tesla.move()
	boing.move()

	myFile := &File{}
	writeToStream(myFile, "hello world")
	readFromStream(myFile)

}
