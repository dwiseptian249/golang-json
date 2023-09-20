package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {

	customer := Customer{
		Fullname: "Dwi Septian",
		Surname:  "Tio",
		City:     "Tangerang",
		Hobbies:  []string{"Gaming", "Coding", "Movie"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"Fullname":"Dwi Septian","Surname":"Tio","City":"Tangerang","Age":0,"Married":false,"Hobbies":["Gaming","Coding","Movie"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Fullname)
	fmt.Println(customer.Hobbies)

}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		Fullname: "Dwi Septian",
		Adresses: []Address{
			{
				Street:     "Jl. Masjid",
				Country:    "Indonesia",
				PostalCode: "585858",
			},
			{
				Street:     "Jl.Lurus",
				Country:    "Indonesia",
				PostalCode: "454545",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"Fullname":"Dwi Septian","Surname":"","City":"","Age":0,"Married":false,"Hobbies":null,"Adresses":[{"Street":"Jl. Masjid","Country":"Indonesia","PostalCode":"585858"},{"Street":"Jl.Lurus","Country":"Indonesia","PostalCode":"454545"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Fullname)
	fmt.Println(customer.Adresses)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jl. Masjid","Country":"Indonesia","PostalCode":"585858"},{"Street":"Jl.Lurus","Country":"Indonesia","PostalCode":"454545"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}

func TestOnlyJSONArrayComplex(t *testing.T) {
	addresses := []Address{
		{
			Street:     "Jl. Masjid",
			Country:    "Indonesia",
			PostalCode: "585858",
		},
		{
			Street:     "Jl.Lurus",
			Country:    "Indonesia",
			PostalCode: "454545",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}
