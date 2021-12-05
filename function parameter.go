package main

import "fmt"

func sayhelloto(Firstname string, Lastname string) {
	fmt.Println("Hello", Firstname, Lastname)
}

func hello() {
	sayhelloto("Eko", "Khaneddy")
}
