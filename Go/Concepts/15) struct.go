package main

import (
)


type UserData struct  { // struct is a collection of fields. 
	fistName string
	lastName string
	email string
	numberOfTickets int
}

func main() {
	user := UserData{ // struct initialization
		fistName: "John",
		lastName: "Doe",
		email: "pradumnasaraf@gmail.com",
		numberOfTickets: 2,
	}

	println(user.fistName) 
}