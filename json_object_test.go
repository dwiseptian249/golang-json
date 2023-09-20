package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

type Customer struct {
	Fullname string
	Surname  string
	City     string
	Age      int
	Married  bool
	Hobbies  []string
	Adresses []Address
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		Fullname: "Dwi Septian",
		Surname:  "Tio",
		City:     "Tangerang",
		Age:      26,
		Married:  false,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
