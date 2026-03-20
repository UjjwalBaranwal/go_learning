package main

import "fmt"

/*
in go
struct → what data looks like + how it behaves
interface → what behavior is expected (contract)
interface is implicit in go
*/
type Person interface {
	GetName() string
	GetId() int
}

type Emp struct {
	ID   int
	Name string
}
type BEmp struct {
	ID   int
	BID  int
	Name string
}

func (e Emp) GetName() string {
	return e.Name
}
func (e Emp) GetId() int {
	return e.ID
}
func (e BEmp) GetName() string {
	return e.Name
}

func (e BEmp) GetId() int {
	return e.ID
}
func display(p Person) {
	fmt.Println(p.GetName(), " ", p.GetId())
}
func main() {
	ujjwal := Emp{
		ID:   1,
		Name: "UJJWAL",
	}
	pro := BEmp{
		ID:   2,
		BID:  1,
		Name: "PRO",
	}
	display(ujjwal)
	display(pro)

	fmt.Println(ujjwal)
}
