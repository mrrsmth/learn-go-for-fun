package filereader

import (
	"fmt"
	"io"
	"os"

)

func Filereader() {
	fmt.Println("init")

	// Открываем файл
	file, err := os.Open("hello2.txt")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}

	// Читаем содержимое файла
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		file.Close()
		return
	}

	// Закрываем файл
	file.Close()

	// Выводим содержимое файла на экран
	fmt.Println(string(content))
}
