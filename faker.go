package faker

import (
	"goFaker/pkg/data/en_GB/bank"
	"goFaker/pkg/data/en_GB/firstnames"
	"goFaker/pkg/data/en_GB/surnames"
	"time"
)

type Person struct {
	Firstname string
	Surname   string
	IBAN      string
}

func GeneratePerson(...interface{}) Person {

	seed := time.Now().UnixNano()

	firstname := firstnames.GenerateFirstname(seed)
	surname := surnames.GenerateSurname(seed)
	IBAN := bank.GenerateIBAN(seed)
	person := Person{firstname, surname, IBAN}

	return person

}
