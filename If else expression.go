package main

import "fmt"

func If_Else_Expression() {
	name := "Joko"
	if name == "Eko" {
		fmt.Println("Hello Eko")
	} else if name == "Joko" {
		fmt.Println("Hi Joko")
	} else {
		fmt.Println("Hi, boleh kenalan?")
	}

	//Short statement if
	var length = len(name)
	if length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
