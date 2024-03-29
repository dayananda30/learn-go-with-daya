package organization

import (
	"errors"
	"fmt"
	"strings"
)

type Identifiable interface {
	//Idetifiable of type interface contains only function names
	ID() string
}

type Person struct {
	// Keeping FirstName and LastName public, anybody can modify them hence we are keeping them as private.
	//FirstName string
	//LastName  string

	firstName      string
	lastName       string
	twitterHandler string
}

func NewPerson(firstName string, lastName string) Person {
	return Person{
		firstName: firstName,
		lastName:  lastName,
	}
}

func (p *Person) FullName() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)

}

func (p *Person) SetTwitterHandler(handler string) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(handler, "@") {
		return errors.New("Twitter Handler must start with @ symbol")
	}

	p.twitterHandler = handler
	return nil

}

func (p *Person) ID() string {
	return "1234"

}
