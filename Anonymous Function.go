package main

import "fmt"

//Anonymous function = function tanpa nama

type Blacklist func(string) bool

func registeruser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

// func blacklistAdmin(name string) bool {
// 	return name == "root"
// }
// func blacklistRoot(name string) bool {
// 	return name == "root"
// }

func Anonymous() {
	blacklist := func(name string) bool {
		return name == "Anjing"
	}

	registeruser("Eko", blacklist)
	registeruser("Anjing", func(name string) bool {
		return name == "Anjing"
	})
}
