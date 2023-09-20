package golangjson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		Fullname: "Tio",
		Surname:  "Septian",
		City:     "Jakarta",
	}

	encoder.Encode(customer)
}
