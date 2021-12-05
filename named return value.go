package main

import "fmt"

func getcompletename() (firstname, middlename, lastname string) {
	firstname = "Eko"
	middlename = "Kurniawan"
	lastname = "Khaneddy"
	return
	//return firstname, middlename, lastname (gk perlu)
}

//cara rename variabel shift + F6 pada sorotan
func named_return_variabel() {
	a, b, c := getcompletename()
	fmt.Println(a, b, c)
}
