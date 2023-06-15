package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	//alex := person{firstName: "Alex", lastName: "Anderson"}
	/*
		var alex person

		// { }
		fmt.Println(alex)

		//{firstName: lastName:}
		fmt.Printf("%+v", alex)

		alex.firstName = "Alex"
		alex.lastName = "Anderson"

		//{Alex Anderson}
		fmt.Println(alex)

		//{firstName:Alex lastName:Anderson}
		fmt.Printf("%+v", alex)
	*/

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 12345,
		},
	}

	//{firstName:Jim lastName:Party contactInfo:{email:jim@gmail.com zipCode:12345}}
	//jim.print()

	// pass by value not working
	//jim.updateNameV1("jimmy")
	//{firstName:Jim lastName:Party contactInfo:{email:jim@gmail.com zipCode:12345}}
	//jim.print()

	jimPointer := &jim
	jimPointer.updateName("jimmy")
	jim.print()

	//Turn address into value with *address
	//Turn value into address with &value

	//SHORTCUT
	jim.updateName("jimmy")
	jim.print()

	name := "Elif"
	//Prints memory address
	fmt.Println(&name)
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

/*func (p person) updateNameV1(newFirstName string) {
	p.firstName = newFirstName
}*/

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
