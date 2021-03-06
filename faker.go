package faker

import (
	"pkg/data/en_GB/automotive"
	"pkg/data/en_GB/bank"
	"pkg/data/en_GB/firstnames"
	"pkg/data/en_GB/server"
	"pkg/data/en_GB/surnames"
	"time"
)

// Bank struct houses the current faker implementations for banking
type Bank struct {
	IBAN          string
	AccountNumber string
	BankCode      string
}

// Names struct houses the current faker implementations for names
type Names struct {
	Firstname string
	Surname   string
}

// Car struct houses the current faker implementations for cars
type Car struct {
	NumberPlate  string
	CurrentSpeed string
}

// Person struct houses the other structs in the base struct "Person"
type Person struct {
	Names Names
	Bank  Bank
	Car   Car
}

// GeneratePerson is the main entry point and populates a
// person object for use globally
func GeneratePerson(...interface{}) Person {

	seed := time.Now().UnixNano()

	firstname := firstnames.GenerateFirstname(seed)
	surname := surnames.GenerateSurname(seed)
	IBAN, AccountNumber, BankCode := bank.GenerateIBAN(seed)
	numberPlate := automotive.GenerateLicensePlate(seed)
	currentSpeed := automotive.GenerateRandomSpeed(seed)
	bankingCredentials := Bank{IBAN, AccountNumber, BankCode}
	names := Names{firstname, surname}
	car := Car{numberPlate, currentSpeed}
	person := Person{names, bankingCredentials, car}

	return person

}

// Server struct houses the other structs in the base struct "Server"
type Server struct {
	File server.File
}

// GenerateServer is a main entry point and populates a
// Server object for use globally
func GenerateServer(...interface{}) Server {

	seed := time.Now().UnixNano()

	fileType := server.GenerateFile(seed)
	serverInstance := Server{fileType}

	return serverInstance
}
