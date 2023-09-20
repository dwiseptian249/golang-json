package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	jsonString := `{"Fullname":"Dwi Septian","Surname":"Tio","City":"Tangerang","Age":26,"Married":false}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Fullname)
	fmt.Println(customer.Surname)
	fmt.Println(customer.City)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
}
