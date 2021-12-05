package main

import "fmt"

/**
STRUCT
1. Template data digunakan untuk menggabungkan nol
atau lebih tipe data lainnya dalam satu kesatuan

2. representasi data dalam program aplikasi yang kita buat

3. data struct disimpan dalam field

4. struct adalah kumpulan field

*/

type Customer struct {
	name, address string
	age           int
}

func Struct() {
	//Membuat data struct
	var eko Customer
	eko.name = "Eko"
	eko.address = "Indonesia"
	eko.age = 30

	fmt.Println(eko.name)
	fmt.Println(eko.address)
	fmt.Println(eko.age)
	// Struct Literals
	joko := Customer{
		name:    "Joko",
		address: "Indonesia",
		age:     30,
	}
	fmt.Println(joko)
	budi := Customer{"Budi", "Indonesia", 6}
	fmt.Println(budi)

}
