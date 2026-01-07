// https://dasarpemrogramangolang.novalagung.com/

package main

import "fmt"

func main() {
	fmt.Println("=== FUNDAMENTAL GO ===")

	// variable
	fmt.Println("Full Name: ", GetFullName())
	fmt.Println("Age Now: ", GetAge())
	// multiple variable
	a, b, c, d := MultiVariable()
	fmt.Println("Multiple variable: ", a, b, c, d)
	// variable underscore
	fmt.Println(UnderscoreVariable())
	// constanta variable
	fmt.Println(ConstantaVariable())
	// constanta multipe variable
	fmt.Println(MyBiodata().FirstName)
	// operator
	OperatorAritmatika()
	OperatorPerbandingan()
	OperatorLogika()
	// control flow
	StrukturKondisional()
	StrukturKondisionalSwitch("senin")
	StrukturPerulangan(100)
	StrukturPerulanganRange()
	// Pointer
	UbahAngkaPointer()
	//  Array Method
	thisIsArray()
	thisIsSlice()
	sliceFromArray()
	thisIsMapMethod()
	// struct
	CreateStruct()
	// method
	PerkenalanMhs()
	// interface
	InterfaceDemo()
}
