package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct {
	Email string
	Phone string
}

type Address struct {
	HouseNo int
	Area    string
	State   string
}

// Nested Struct
type Employee struct {
	Personal_Details Person
	Person_Contact   Contact
	Personal_Address   Address
}

func main() {
	var gaurav Person

	gaurav.FirstName = "Gaurav"
	gaurav.LastName = "Subedi"
	gaurav.Age = 23

	fmt.Println("Gaurav's full details are:", gaurav)

	// We can also use the shorthand way to create a struct
	subigya := Person{
		FirstName: "Subigya",
		LastName:  "Subedi",
		Age:       22,
	}

	fmt.Println("Subigya's full details are:", subigya)

	// Using new keyword
	var ronaldo = new(Person)

	ronaldo.FirstName = "Cristiano"
	ronaldo.LastName = "Ronaldo"
	ronaldo.Age = 40
	fmt.Println("Ronaldo's full details are:", ronaldo)

	fmt.Println("First Name of gaurav is:", gaurav.FirstName)
	fmt.Println("Last Name of subigya is:", subigya.LastName)
	fmt.Println("Age of ronaldo is:", ronaldo.Age)

	// Nested Struct
	employee1 := Employee{
		Personal_Details: Person{
			FirstName: "Gaurav",
			LastName:  "Subedi",
			Age:       23,
		},

		Person_Contact: Contact{
			Email: "gaurav@gmail.com",
			Phone: "1234567890",
		},

		Personal_Address: Address {
			HouseNo: 1,
			Area:    "Kathmandu",
			State:   "Nepal",
		},
	}

	fmt.Println("Employee 1 details are:", employee1)
	fmt.Println("Employee 1's First Name is:", employee1.Personal_Details)
	fmt.Println("Employee 1's Email is:", employee1.Person_Contact.Email)
}
