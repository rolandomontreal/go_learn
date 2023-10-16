package main

import "fmt"

type person struct {
  name string
  contactInfo
}

type contactInfo struct {
  email string
}

func main() {
  alex := person{ 
    firstName: "Alex", 
    lastName: "Anderson", 
    contactInfo: contactInfo{
      email: "mail@mail.com",
      zipCode: 12323,
    },
  }

  alex.updateName("Steve")

  alex.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
  pointerToPerson.firstName = newFirstName
}

func (p person) print() {
  fmt.Printf("The person: %+v \n\n", p)
}