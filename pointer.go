package main

import "fmt"

func UbahAngka(x *int) {
	*x = 100
}

func UbahAngkaPointer() {
	a := 10
	UbahAngka(&a)
	fmt.Println(a)
}
