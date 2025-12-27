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
	return Person{
		FirstName:  DefaultFirstName,
		LastName:   DefaultLastName,
		Age:        DefaultAge,
		Address:    DefaultAddress,
		HomeNumber: DefaultHomeNumber,
	}
}
