package main

import (
	"fmt"
	"time"
)

type Emp struct {
	ID        int
	FirstName string
	LastName  string
	Position  string
	Salary    int
	IsActve   bool
	JoinedAt  time.Time
}

func createNewEmp(id int, firstName string, lastName string, position string, salary int, isActive bool, joinedAt time.Time) Emp {
	return Emp{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Position:  position,
		Salary:    salary,
		IsActve:   isActive,
		JoinedAt:  joinedAt,
	}
}

func (e Emp) FullName() string {
	return e.FirstName + " " + e.LastName + "\n"
}

func (e *Emp) Deactivate() {
	e.IsActve = false
}

func (e *Emp) SetJoinedAt(t time.Time) {
	e.JoinedAt = t
}
func main() {
	ujjwal := Emp{
		ID:        1,
		FirstName: "Ujjwal",
		LastName:  "Baranwal",
		Position:  "ADMIN",
		Salary:    10,
		IsActve:   true,
		JoinedAt:  time.Now(),
	}
	fmt.Printf("%+v\n", ujjwal)
	fmt.Println(ujjwal.FirstName)

	pro := createNewEmp(2, "Pro", "Pro", "USER", 100, true, time.Now())
	fmt.Printf("%+v\n", pro)

	proPtr := &pro
	proPtr.FirstName = "Proooo"
	fmt.Printf("%+v\n", pro)
	fmt.Println(pro.FullName())
	pro.Deactivate()
	fmt.Printf("%+v\n", pro)
	pro.SetJoinedAt(time.Now().Add(100000000 * time.Minute))
	fmt.Printf("%+v\n", pro)

	// joe := &Emp{}
	// joe = nil // demo to show panic

	// joe.FullName() // Bad idea
}
