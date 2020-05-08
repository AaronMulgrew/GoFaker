# GoFaker
[![Build Status](https://travis-ci.org/AaronMulgrew/GoFaker.svg?branch=master)](https://travis-ci.org/AaronMulgrew/GoFaker) [![codecov](https://codecov.io/gh/AaronMulgrew/GoFaker/branch/master/graph/badge.svg)](https://codecov.io/gh/AaronMulgrew/GoFaker) [![Go Report Card](https://goreportcard.com/badge/github.com/aaronmulgrew/gofaker)](https://goreportcard.com/report/github.com/aaronmulgrew/gofaker)

Faker is a Go package to generate fake data. It can be used for realistic test data. Inspired by [Python Faker](https://github.com/joke2k/faker)

# Why use this library?
There are many faker libraries out there already. However, none of them are comprehensive, performant and easy to use. The aim of this library is to combine those things into a quick, easy to use and extensible library, using only core Golang libraries to generate the data. 

# Installation
```cmd
go get -u github.com/AaronMulgrew/GoFaker
```

# Usage
Simply import the library and generate a Person struct to accessible available data.
```Golang
package main

import (
  "fmt"
  "github.com/AaronMulgrew/GoFaker"
)

func main() {
  person := faker.GeneratePerson()
  fmt.Println(person.Names.Firstname)
  fmt.Println(person)
}
```


You can also view the [Faker Client](https://github.com/AaronMulgrew/FakerClient) for a full list of available structs and other examples.

# Currently Supported Data
- Banking
  - IBAN (Should validate against IBAN validators)
  - Account Number
  - Bank Code (Does not use real bank codes)
- Automotive
  - Number Plates
- Names
  - Firstname
  - Surname

# Currently Supported Regional Data
- English (United Kingdom)
