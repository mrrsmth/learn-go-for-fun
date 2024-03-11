package codewars

import "fmt"

var maps = []int{1, 2, 3}

func Hello() {
	fmt.Println("Привет из Codewars!")
	arr := []int{1, 2, 3}
	result := Maps(arr)
	fmt.Println(result)
}

// Given an array of integers, return a new array with each value doubled.

// For example:

// [1, 2, 3] --> [2, 4, 6]

func Maps(x []int) []int {
	newX := make([]int, len(x))

	for _, value := range x {
		doubledValue := value * 2
		newX = append(newX, doubledValue)
	}

	return newX
}
