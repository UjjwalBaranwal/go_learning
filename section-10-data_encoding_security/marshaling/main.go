package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
	Marshaling

Converting a Go object → into a transferable format (like JSON)
Go struct → JSON string
*/
// type User struct {
// 	Name  string
// 	Age   int
// 	Phone int
// 	// isActive bool // this will not work
// 	// JSON only works with capitalized (exported) fields
// 	IsActive bool
// }

// if we want as encoded json name different from given
type User struct {
	Name  string `json:"full_name"`
	Age   int    `json:"age"`
	Phone int    `json:"phone_number"`
	// isActive bool // this will not work
	// JSON only works with capitalized (exported) fields
	IsActive bool `json:"is_active"`
}

func main() {
	ujjwal := User{
		Name:     "Ujjwal",
		Age:      23,
		Phone:    1234567890,
		IsActive: true,
	}
	// byteSlice, err := json.Marshal(ujjwal)
	// byteSlice, err := json.MarshalIndent(ujjwal, "json-", "-json")
	byteSlice, err := json.MarshalIndent(ujjwal, "", "")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(byteSlice))
}
