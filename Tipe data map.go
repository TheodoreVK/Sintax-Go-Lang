package main

import "fmt"

func Tipe_Data_Map() {
	person := map[string]string{
		"name":    "Eko",
		"address": "Subang",
	}

	fmt.Println(person)
	fmt.Println(person["name"]) //mengambil data di map dengan key
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] = "buku Go-Lang"
	book["author"] = "Eko Kurniawan"
	book["wrong"] = "ups"
	fmt.Println(book)

	delete(book, "wrong")

	fmt.Println(book)
}
