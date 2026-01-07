package main

import "fmt"

// array method
func thisIsArray() {
	numbers := [5]int{10, 20, 30, 40, 50}

	fmt.Println("Array: ", numbers)
}

// slice method
func thisIsSlice() {
	numbers := []int{10, 20, 30, 40, 50}

	addNumber := append(numbers, 60)

	fmt.Println("Array Slice: ", addNumber)
}

// buat slice dari array
func sliceFromArray() {
	numbers := [5]int{10, 20, 30, 40, 50}
	sliceNumbers := numbers[1:4]

	fmt.Println("Slice from Array: ", sliceNumbers)
}

// map method
func thisIsMapMethod() {
	nilaiMap := map[string]int{
		"Andi": 90,
		"Budi": 85,
	}

	nilaiMap["Caca"] = 95

	fmt.Println("Map Data: ", nilaiMap)

	fmt.Println("Nilai Caca: ", nilaiMap["Caca"])
}
