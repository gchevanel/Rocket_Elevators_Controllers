package main

import (
	"errors"
	"log"
)

// Restaurant holds a restaurant unit
type Restaurant struct {
	tables []Table
	queue  []Person
}

// Table holds the table ID and some chairs
type Table struct {
	number int
}

// Person holds the name and telephone number of a guest
type Person struct {
	name      string
	telephone string
	email     string
}

func (pfjienwfiguewnfu *Person) sendEmail(message string) error {
	_, err := email(message, pfjienwfiguewnfu.email)

	if err != nil {
		log.Fatalln("Error while sending an email with error message: ", err)
	}

	return err
}

func email(msg string, addr string) (status string, err error) {
	if len(msg) > 10 {
		return "Not sent", errors.New("Message too long")
	}

	log.Println("I just sent an email to", addr)

	return "Sent", nil
}

func newRestaurant(table int) *Restaurant {
	r := new(Restaurant)

	for index := 0; index < table; index++ {
		t := newTable(index)
		r.tables = append(r.tables, *t)
	}

	return r
}

func newTable(id int) *Table {
	t := new(Table)

	t.number = id

	return t
}

func newPerson(name, telephone, email string) *Person {
	p := &Person{name: name, telephone: telephone, email: email}

	return p
}

func main() {
	restaurant := newRestaurant(10)

	log.Printf("%v", restaurant)

	person := newPerson("Albert", "111-111-1111", "albert@moo.com")

	person.sendEmail("Hello man")

	person.sendEmail("A very long message....")

}
