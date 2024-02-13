package stream

import "fmt"

// type Stream interface {
// 	read() string
// 	write(string)
// 	close()
// }

// func writeToStream(stream Stream, text string) {
// 	stream.write(text)
// }
// func closeStream(stream Stream) {
// 	stream.close()
// }

// // структура файл
// type File struct {
// 	text string
// }

// // структура папка
// type Folder struct{}

// // реализация методов для типа *File
// func (f *File) read() string {
// 	return f.text
// }
// func (f *File) write(message string) {
// 	f.text = message
// 	fmt.Println("Запись в файл строки", message)
// }
// func (f *File) close() {
// 	fmt.Println("Файл закрыт")
// }

// // релизация методов для типа *Folder
// func (f *Folder) close() {
// 	fmt.Println("Папка закрыта")
// }

// type Reader interface {
// 	read()
// }

// type Writer interface {
// 	write(string)
// }

// func writeToStream(writer Writer, text string) {
// 	writer.write(text)
// }
// func readFromStream(reader Reader) {
// 	reader.read()
// }

// type File struct {
// 	text string
// }

// func (f *File) read() {
// 	fmt.Println(f.text)
// }
// func (f *File) write(message string) {
// 	f.text = message
// 	fmt.Println("Запись в файл строки", message)
// }

type Reader interface {
	read()
}

type Writer interface {
	write(string)
}

type ReaderWriter interface {
	Reader
	Writer
}

func writeToStream(writer ReaderWriter, text string) {
	writer.write(text)
}
func readFromStream(reader ReaderWriter) {
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

func StreamMain() {

	// myFile := &File{}
	// myFolder := &Folder{}

	// writeToStream(myFile, "hello world")
	// closeStream(myFile)
	// //closeStream(myFolder)     // Ошибка: тип *Folder не реализует интерфейс Stream
	// myFolder.close() // Так можно

	// myFile := &File{}
	// writeToStream(myFile, "hello world")
	// readFromStream(myFile)

	myFile := &File{}
	writeToStream(myFile, "hello world")
	readFromStream(myFile)
	writeToStream(myFile, "lolly bomb")
	readFromStream(myFile)
}
