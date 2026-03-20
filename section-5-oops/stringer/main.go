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
func (e BEmp) String() string {
	// this format will use when using fmt.(PrintStatemnet)
	return fmt.Sprintf("BEMP [id : %d,bid : %d,name : %s]", e.ID, e.BID, e.Name)
}

func display(p Person) {
	fmt.Println(p.GetName(), " ", p.GetId())
}

type Id int

func (id Id) String() string {
	return fmt.Sprintf("Id : %d", id)
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
	fmt.Println(pro)

	fmt.Println(ujjwal)
	id := Id(30)
	fmt.Println(id)
}
