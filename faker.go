package faker


import (
	"goFaker/pkg/data/en_GB/firstnames"
	"goFaker/pkg/data/en_GB/surnames"
	"time"
)

type Person struct {
	Firstname string
	Surname string
}




func GeneratePerson(...interface{}) Person {

	seed := time.Now().UnixNano()
	
	firstname := firstnames.GenerateFirstname(seed);
	surname := surnames.GenerateSurname(seed);

	person := Person{firstname, surname};
	
	return person

}