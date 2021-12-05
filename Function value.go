package main

import "fmt"

func getgoodbye(name string) string {
	return "Good Bye" + name
}

func goodbye() {
	goodbye := getgoodbye
	fmt.Println(goodbye(" Eko"))
}
