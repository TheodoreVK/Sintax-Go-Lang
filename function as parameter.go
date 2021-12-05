package main

import "fmt"

//type Filter func(string) string
// opsional, jika dipakai maka tulisan cukup seperti ini
//func sayhellowithfilter (name(string), filter Filter)

func sayhellowithfilter(name string, filter func(string) string) {
	fmt.Println("Hello", filter(name))
}

func spamfilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func spam() {
	sayhellowithfilter("Eko", spamfilter)
	sayhellowithfilter("Anjing", spamfilter)
}
