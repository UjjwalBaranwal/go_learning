package main

import (
	"fmt"
	"strings"
)

func main() {
	// s1 := "abc"
	// s2 := s1
	// fmt.Printf("Address of s1 : %x, address of s2 : %x\n", &s1, &s2)
	// s2 = "bcd"
	// fmt.Printf("s1 : %s, s2: %s\n", s1, s2)
	s1 := "abc"
	s2 := strings.Clone(s1)

	fmt.Println(s1, s2)
	b := strings.Builder{}
	b.WriteString("Ujjwal")
	b.WriteString("Baranwal")
	b.WriteRune('世')
	fmt.Println(b.String())

	fmt.Println(strings.ToLower(b.String()))
	fmt.Println(strings.ToUpper(s1))
	fmt.Println(strings.Title(s1))
	s3 := "     test sss    "
	fmt.Println("s3", len(s3))
	s3 = strings.TrimSpace(s3)

	fmt.Println("s3 after trim", len(s3))
	fmt.Println(strings.HasSuffix("test@gmail.com", "gmail.com"))
	fmt.Println(strings.HasPrefix("test@gmail.com", "test"))
	fmt.Println(strings.Replace("test@gmail.com", "test", "john", 1))
	fmt.Println(strings.Replace("test@gmailtest.com", "test", "john", 2))
	parts := strings.Split("test@gmail.com", "@")
	username, domain := parts[0], parts[1]
	fmt.Println(username, domain)

	parts = strings.Fields("jane example.com")
	username, domain = parts[0], parts[1]
	fmt.Println(username, domain)

	fmt.Println(strings.Join(parts, ","))
}
