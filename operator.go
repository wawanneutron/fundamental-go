package main

import "fmt"

func OperatorAritmatika() {
	a := 10
	b := 20

	fmt.Println("penjumlahan: ", a+b)
	fmt.Println("pengurangan: ", b-a)
	fmt.Println("perkalian: ", a*b)
	fmt.Println("pembagian: ", b/a)
	fmt.Println("modules sisa bagi: ", a%b)
}

func OperatorPerbandingan() {
	x := 10
	y := 20

	fmt.Println("lebih kecil: ", x < y)              // true
	fmt.Println("lebih besar: ", x > y)              // false
	fmt.Println("sama dengan: ", x == y)             // false
	fmt.Println("tidak sama dengan: ", x != y)       // true
	fmt.Println("lebih kecil sama dengan: ", x <= y) // true
	fmt.Println("lebih besar sama dengan: ", x >= y) // false
}

func OperatorLogika() {
	a := true
	b := false

	// AND
	fmt.Println("AND: ", a && b) // false
	// OR
	fmt.Println("OR: ", a || b) // true
	// NOT
	fmt.Println("NOT a: ", !a) // false
	fmt.Println("NOT b: ", !b) // true
}
