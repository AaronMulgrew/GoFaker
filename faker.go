package faker


import (
    "fmt"
	"goFaker/pkg/data/names"
	"time"
)



func AllUsers(...interface{}) {

	seed := time.Now().UnixNano()
	
	name := names.Names(seed);
	fmt.Println(name);
	fmt.Println("Hello!");

}