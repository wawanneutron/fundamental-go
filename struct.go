package main

import "fmt"

type Alamat struct {
	Kota  string
	Jalan string
	Nomor int
}

type Kampus struct {
	Nama   string
	Ijazah bool
}

type Mahasiswa struct {
	Nama   string
	Umur   int
	Aktif  bool
	Alamat Alamat
	Kampus Kampus
}

// pointer ke struct
func UbahUmur(mhs *Mahasiswa, umurBaru int) {
	mhs.Umur = umurBaru
}

func CreateStruct() {
	// create variable dari struct
	mhs1 := Mahasiswa{}

	mhs1.Nama = "Jhon Dhoe"
	mhs1.Umur = 22
	mhs1.Aktif = true
	mhs1.Alamat = Alamat{
		Kota:  "Jakarta",
		Jalan: "JL Jendral Sudirman",
		Nomor: 20,
	}
	mhs1.Kampus = Kampus{
		Nama:   "UI",
		Ijazah: false,
	}

	// struct literal (bisa langsung assign)
	mhs2 := Mahasiswa{
		Nama:  "Jokowi",
		Umur:  23,
		Aktif: false,
		Alamat: Alamat{
			Kota:  "Solo",
			Jalan: "Jalan Merdeka",
			Nomor: 36,
		},
		Kampus: Kampus{
			Nama:   "UGM",
			Ijazah: false,
		},
	}

	mhs3 := Mahasiswa{
		Nama:  "Wawan Setiawan",
		Umur:  25,
		Aktif: false,
		Alamat: Alamat{
			Kota:  "Tangerang",
			Jalan: "Jalan Dewantara",
			Nomor: 77,
		},
		Kampus: Kampus{
			Nama:   "ITB",
			Ijazah: true,
		},
	}

	// pointer ke struct
	UbahUmur(&mhs1, 30)
	UbahUmur(&mhs2, 53)
	UbahUmur(&mhs3, 27)

	isMhsAktif := mhs1.Aktif || mhs2.Aktif || mhs3.Aktif
	isAnyIjazah := mhs1.Kampus.Ijazah || mhs2.Kampus.Ijazah || mhs3.Kampus.Ijazah

	if isMhsAktif && isAnyIjazah {
		if mhs2.Nama == "Jokowi" && mhs3.Kampus.Ijazah {
			fmt.Println("Ijazah Jokowi Palsu")
		}

		fmt.Println("Ada Mahasiswa Aktif dan Punya Ijazah")
	}

	fmt.Println("Is Mahasiswa Aktif: ", isMhsAktif)
	fmt.Println("Is Any Mahasiswa Ijazah: ", isAnyIjazah)

	fmt.Println("Mahasiswa 1: ", mhs1)
	fmt.Println("Mahasiswa 2: ", mhs2)
	fmt.Println("Mahasiswa 3: ", mhs3)
	fmt.Println("Akses Field Mahasiswa 2: ", mhs2.Nama)
}
