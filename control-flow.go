package main

import (
	"fmt"
	"strings"
)

// struktur kondisional if else
func StrukturKondisional() {
	nilai := 55

	if nilai >= 85 {
		fmt.Println("Nilai A")
	} else if nilai >= 70 {
		fmt.Println("Nilai B")
	} else if nilai >= 60 {
		fmt.Println("Nilai C")
	} else {
		fmt.Println("Nilai D")
	}
}

// switch case
func StrukturKondisionalSwitch(hari string) {
	days := strings.ToLower(hari)

	switch days {
	case "senin":
		fmt.Println("Hari Senin awal minggu")
	case "selasa":
		fmt.Println("Hari kedua")
	case "rabu":
		fmt.Println("Hari ketiga")
	case "kamis":
		fmt.Println("Hari keempat")
	case "jumat":
		fmt.Println("Hari kelima, menjelang weekend")
	case "sabtu", "minggu":
		fmt.Println("Hari Weekend, yeayyy")
	default:
		fmt.Println("inputan tidak sesuai")
	}
}

// perulangan biasa
func StrukturPerulangan(value int) {
	for i := 1; i <= value; i++ {
		genap := i % 20

		fmt.Println("Perulangan: ", genap)

		if genap == 0 {
			fmt.Println("Perulangan genap: ", i)
		}
	}
}

// perulangan dengan range
func StrukturPerulanganRange() {
	slice := []string{"Hendrawan", "Wawan", "Setiawan"}

	for index, value := range slice {
		if index == 1 {
			fmt.Println("Hallo nama saya ", value)
		}

		fmt.Println("Index: ", index, "Value: ", value)
	}
}
