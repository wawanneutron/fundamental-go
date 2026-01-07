package main

// Multipe Constant

const (
	DefaultFirstName  = "Wawan"
	DefaultLastName   = "Setiawan"
	DefaultAge        = 28
	DefaultAddress    = "Tangerang"
	DefaultHomeNumber = 36
)

type Person struct {
	FirstName  string
	LastName   string
	Age        int
	Address    string
	HomeNumber int
}

func MyBiodata() Person {
	human := Person{}
	human.FirstName = "Jokowi"

	return Person{
		FirstName:  human.FirstName,
		LastName:   DefaultLastName,
		Age:        DefaultAge,
		Address:    DefaultAddress,
		HomeNumber: DefaultHomeNumber,
	}
}
