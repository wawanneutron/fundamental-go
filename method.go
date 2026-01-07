package main

import "fmt"

type MahasiswaBaru struct {
	Nama       string
	Umur       int
	TahunLahir int
}

func (m MahasiswaBaru) Perkenalan() {
	fmt.Printf("Halo, nama saya %s umur saya %d dan tahun lahir saya %d\n", m.Nama, m.Umur, m.TahunLahir)
}

// method dengan return value
func (m MahasiswaBaru) HitungTahunLahir(tahunSekarang int) int {
	return tahunSekarang - m.Umur
}

// update dengan pointer
func UpdateTahunLahir(m *MahasiswaBaru, tahunLahirBaru int) {
	m.TahunLahir = tahunLahirBaru
}

// jalankan untuk di main go
func PerkenalanMhs() {

	mhs := MahasiswaBaru{
		Nama: "Jokowi",
		Umur: 29,
	}

	thnLahir := mhs.HitungTahunLahir(2026)
	UpdateTahunLahir(&mhs, thnLahir)

	mhs.Perkenalan()
}
