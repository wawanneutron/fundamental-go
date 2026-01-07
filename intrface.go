package main

import "fmt"

type BisaBerjalan interface {
	Jalan()
}

type Manusia struct {
	Nama string
}

type Robot struct {
	Kode string
}

func (m Manusia) Jalan() {
	fmt.Println(m.Nama, "sedang berjalan")
}

func (r Robot) Jalan() {
	fmt.Println("Robot", r.Kode, "Berjalan otomatis")
}

func Gerakan(b BisaBerjalan) {
	b.Jalan()
}

func InterfaceDemo() {
	manusia := Manusia{Nama: "Jokowi"}
	robot := Robot{Kode: "RBT-804"}

	Gerakan(manusia)
	Gerakan(robot)
}
