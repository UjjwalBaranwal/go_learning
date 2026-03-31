package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	Name     string `json:"name" xml:"name"`
	Age      int    `json:"age" xml:"age"`
	Phone    int    `json:"phone" xml:"phone_number"`
	Password string `json:"-" xml:"-"`
	IsActive bool   `json:"active" xml:"active"`
}

func main() {
	u := user{
		Name:  "Ujjwal",
		Age:   30,
		Phone: 9012893490,
	}

	// enc := json.NewEncoder(os.Stdout)
	buffer := new(bytes.Buffer)
	enc := json.NewEncoder(buffer)
	if err := enc.Encode(&u); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buffer)
}
