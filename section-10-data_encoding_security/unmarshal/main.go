package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
	UnMarshal

Converting data (JSON) -> back to Go object
JSON string → Go struct
*/
type User struct {
	Name  string `json:"full_name"`
	Age   int    `json:"age"`
	Phone int    `json:"phone_number"`
	// isActive bool // this will not work
	// JSON only works with capitalized (exported) fields
	IsActive bool    `json:"is_active"`
	Profile  profile `json:"profile"`
}

type profile struct {
	URL string `json:"url"`
}

// var payload = `{"full_name":"Ujjwal","age":23,"phone_number":1234567890,"is_active":true}`

var payload = `{
 "full_name": "Ujjwal",
 "age": 20,
 "phone_number": 123456789,
 "is_active": true,
 "profile": {
  "url": "https://www.ujjwal.co.id"
 }
}
`

func main() {
	var u User
	err := json.Unmarshal([]byte(payload), &u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", u)
	fmt.Println(u.Age)
	bs, err := json.MarshalIndent(u, "", "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bs))
}
