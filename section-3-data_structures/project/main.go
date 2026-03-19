package main

import "fmt"

//simple contact management system

type Contact struct{
	Id int
	Name string
	Email string
	Phone string
}

var contactList []Contact
var contactIndexByName map[string]int
var nextId =1

func init(){
	contactList=make([]Contact,0)
	contactIndexByName=make(map[string]int)
}

func addContactToList(name string,email string, phone string){
	if _,exist:=contactIndexByName[name];exist{
		fmt.Println("Contact already exist , name : ",name)
		return
	}
	newContact :=Contact{
		Id: nextId,
		Name: name,
		Email: email,
		Phone: phone,
	}
	nextId++;
	contactList=append(contactList, newContact)
	contactIndexByName[name]=len(contactList)-1;
	fmt.Println("Contact is added , name : ",name)
}

func findContact(name string) *Contact{
	ind,exist:=contactIndexByName[name]
	if exist{
		return &contactList[ind]
	}
	return nil
}

func listAllContact(){
	fmt.Println("Listing all contact")

	if len(contactList) == 0{
		fmt.Println("Contact List is empty")
	}else{
		for _,contact := range contactList{
			fmt.Printf("{id : %d , name : %s , email : %s , phone : %s}\n",contact.Id,contact.Name,contact.Email,contact.Phone)
		}
	}

}

func main(){
	addContactToList("Ujjwal","u@gmail.com","909090")
	addContactToList("UjjwalB","ub@gmail.com","909091")
	addContactToList("UjjwalBU","ubu@gmail.com","909092")
	addContactToList("UjjwalBU","ubu@gmail.com","909092")

	listAllContact()

	ujjwal:=findContact("Ujjwal")
	if ujjwal != nil{
		fmt.Println("Ujjwal is in contact list")
	}else{
		fmt.Println("Ujjwal not in contact list")
	}

	pro:=findContact("Pro")
	if pro != nil{
		fmt.Println("Pro is in contact list")
	}else{
		fmt.Println("Pro is not in contact list")
	}
}