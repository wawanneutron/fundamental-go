package main

import "time"

func GetFullName() string {
	firstName := "Wawan"
	lastName := "Hendrawan"

	lastName = "Setiawan"

	return firstName + " " + lastName
}

func GetAge() int {
	year := 1997
	currentYear := time.Now().Year()

	return currentYear - year
}

// multiple variable
func MultiVariable() (int, bool, float64, string) {
	one, isSunday, twoPointTwo, say := 1, true, 2.2, "Hello"

	return one, isSunday, twoPointTwo, say
}

// variable Underscore
func UnderscoreVariable() string {
	name, _ := "Jhon", "Dhoe"

	return "My name is" + " " + name
}

// constanta variable
func ConstantaVariable() string {
	const firstName = "Wawan"
	const lastName = "Setiawan"

	// lastName = "Hendrawan" cannot assign data type constanta

	return firstName + " " + lastName
}

type Biodata struct {
	name    string
	age     int
	Address string
}
