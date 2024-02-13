package file

import (
	"fmt"
	"io"
	"os"

)

type phoneReader string

func (ph phoneReader) Read(p []byte) (int, error) {
	count := 0
	for i := 0; i < len(ph); i++ {
		if ph[i] >= '0' && ph[i] <= '9' {
			p[count] = ph[i]
			count++
		}
	}
	return count, io.EOF
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type person struct {
	name   string
	age    int32
	weight float64
}

func writeData(filename string) {
	// начальные данные
	var name string = "Tom"
	var age int = 24
	var text string = "lorem10"

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Fprintln(file, name, text)
	fmt.Fprintln(file, age, text)
}
func readData(filename string) {

	var name string
	var age int
	var text string
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Fscanln(file, &name)
	fmt.Fscanln(file, &age)
	fmt.Fscanln(file, &text)
	fmt.Println(name, age, text)
}

// func writeData(filename string) {
// 	// начальные данные
// 	tom := person{name: "Tom", age: 24, weight: 68.5}

// 	file, err := os.Create(filename)
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	defer file.Close()

// 	// сохраняем данные в файл
// 	fmt.Fprintf(file, "%s %d %.2f\n", tom.name, tom.age, tom.weight)
// }
// func readData(filename string) {

// 	// переменные для считывания данных
// 	var name string
// 	var age int
// 	var weight float64

// 	file, err := os.Open(filename)
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	defer file.Close()

// 	// считывание данных из файла
// 	_, err = fmt.Fscanf(file, "%s %d %f\n", &name, &age, &weight)
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	fmt.Printf("%-8s %-8d %-8.2f\n", name, age, weight)
// }

func FileMain() {

	// 	filename := "person.dat"
	// 	writeData(filename)
	// 	readData(filename)

	filename := "hello2.txt"
	writeData(filename)
	readData(filename)

	// fmt.Fprintln(os.Stdout, "hello cold")

	// tom := person{
	// 	name:   "Tom",
	// 	age:    24,
	// 	weight: 68.5,
	// }
	// file, err := os.Create("person.dat")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// defer file.Close()
	// fmt.Fprintf(file,
	// 	"%-10s %-10d %-10.3f\n",
	// 	tom.name, tom.age, tom.weight)

	// phone1 := phoneReader("+1(234)567 90-10")
	// io.Copy(os.Stdout, phone1)
	// fmt.Println()

	// file, err := os.Create("confeve.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// defer file.Close()
	// fmt.Fprint(file, "Сегодня ")
	// fmt.Fprintln(file, "хорошая погода")

	// file, err := os.Open("hello.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// defer file.Close()

	// io.Copy(os.Stdout, file)

	// dat, err := os.ReadFile("hello.txt")
	// check(err)
	// fmt.Print(string(dat))

	// for _, value := range dat {
	// 	fmt.Println(string(value))
	// }
	// fmt.Println("init File")
	// phone1 := phoneReader("+1(234)567 90104")
	// phone2 := phoneReader("+2-345-678-12-35")

	// buffer := make([]byte, len(phone1))
	// phone1.Read(buffer)
	// fmt.Println(string(buffer)) // 12345679010

	// buffer = make([]byte, len(phone2))
	// phone2.Read(buffer)
	// fmt.Println(string(buffer)) // 23456781235

	// text := "console.log(1)"
	// file, err := os.Create("hello.js")

	// if err != nil {
	// 	fmt.Println("Unable to create file:", err)
	// 	os.Exit(1)
	// }
	// defer file.Close()
	// file.WriteString(text)

	// fmt.Println("Done.")

	// data := []byte("Hello Bold!")
	// file, err := os.Create("hello.bin")
	// if err != nil {
	// 	fmt.Println("Unable to create file:", err)
	// 	os.Exit(1)
	// }
	// defer file.Close()
	// file.Write(data)

	// fmt.Println("Done.")

	// file, err := os.Open("hello.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// defer file.Close()

	// data := make([]byte, 64)

	// for {
	// n, err := file.Read(data)
	// 	if err == io.EOF { // если конец файла
	// 		break // выходим из цикла
	// 	}
	// 	fmt.Print(string(data[:n]))
	// }

}
