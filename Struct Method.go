package main

import "fmt"

// type Customer struct {
// 	name, address string
// 	age           int
// }

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello, My Name is", customer.name)
}

func Struct_Method() {
	//Membuat data struct
	var eko Customer
	eko.name = "Eko"
	eko.address = "Indonesia"
	eko.age = 30

	eko.sayHi("Joko")

	// fmt.Println(eko.name)
	// fmt.Println(eko.address)
	// fmt.Println(eko.age)
	// // Struct Literals
	// joko := Customer{
	// 	name:    "Joko",
	// 	address: "Indonesia",
	// 	age:     30,
	// }
	// fmt.Println(joko)
	// budi := Customer{"Budi", "Indonesia", 6}
	// fmt.Println(budi)

}
