package main

import "fmt"

func main() {
	fmt.Println("Hello Go!")
	fmt.Println("I am learnign go for backend developement.")
	/*
		name := "Amin"
		fmt.Println(name)

		name = "Dani"
		fmt.Println(name)

		var userName string = "This is my userName!"
		println(userName)

		var password string
		var age int
		var price float32
		var isDelete bool

		fmt.Println(password)
		fmt.Println(age)
		fmt.Println(price)
		fmt.Println(isDelete)

		var floatTest float32 = 200.12
		changingType := int(floatTest)
		println(changingType)
	*/

	//Practice 002
	//01:
	Name := "AminG82"
	Age := 22
	Email := "AminG82@gmail.com"
	IsMember := false

	println(Name)
	println(Age)
	println(Email)
	println(IsMember)
	//02:
	var name string
	var age int
	var price float64
	var isMember bool

	println(name)
	println(age)
	println(price)
	println(isMember)
	//03:
	ticketPrice := 25000.0
	quantity := 4

	total := ticketPrice * float64(quantity)

	println(total)

}
