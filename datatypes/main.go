package main

import (
	"dayananda30/learn-go-with-daya/datatypes/organization"
	"fmt"
)

func main() {
	//p := organization.Person{FirstName: "Dayananda", LastName: "D R"}

	p := organization.NewPerson("Dayananda", "D R")

	err := p.SetTwitterHandler("@twitter.com")

	if err != nil {
		fmt.Printf("An Error Encountered - %s \n", err.Error())

	}

	fmt.Println(p.ID())
	fmt.Println(p.FullName())
}
