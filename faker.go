package faker

import (
	"goFaker/pkg/data/en_GB/automotive"
	"goFaker/pkg/data/en_GB/bank"
	"goFaker/pkg/data/en_GB/firstnames"
	"goFaker/pkg/data/en_GB/surnames"
	"time"
)

type Bank struct {
	IBAN          string
	AccountNumber string
	BankCode      string
}

type Names struct {
	Firstname string
	Surname   string
}

type Car struct {
	NumberPlate string
}

type Person struct {
	Names Names
	Bank  Bank
	Car   Car
}

func GeneratePerson(...interface{}) Person {

	seed := time.Now().UnixNano()

	firstname := firstnames.GenerateFirstname(seed)
	surname := surnames.GenerateSurname(seed)
	IBAN, AccountNumber, BankCode := bank.GenerateIBAN(seed)
	numberPlate := automotive.GenerateLicensePlate(seed)
	bankingCredentials := Bank{IBAN, AccountNumber, BankCode}
	names := Names{firstname, surname}
	car := Car{numberPlate}
	person := Person{names, bankingCredentials, car}

	return person

}
