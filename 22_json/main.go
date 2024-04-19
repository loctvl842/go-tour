package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Define a JSON string
	jsonString := `{"name": "John", "age": 30, "city": "New York"}`

	// Convert the JSON string to a byte slice
	jsonBytes := []byte(jsonString)

	// Define a struct to unmarshal the JSON data into
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		City string `json:"city"`
	}

	// Unmarshal the JSON byte slice into the struct
	var person Person
	err := json.Unmarshal(jsonBytes, &person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the unmarshaled struct
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
	fmt.Println("City:", person.City)
}
