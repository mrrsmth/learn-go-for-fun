package math

import (
	"fmt"
	"strconv"
)

func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	println(total)
	println(total / float64(len(xs)))
	return total / float64(len(xs))
}

func StrSlice() {
	strSlice := []string{"1.3", "2"}
	numbers := make([]float64, len(strSlice))

	for i, str := range strSlice {
		num, err := strconv.ParseFloat(str, 64)
		if err != nil {
			fmt.Printf("Ошибка парсинга значения '%s': %s\n", str, err.Error())
			continue
		}
		numbers[i] = num
	}

	fmt.Println("Преобразованный срез:", numbers)
}
