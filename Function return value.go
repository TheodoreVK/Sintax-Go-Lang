package main

import "fmt"

func gethello(name string) string {
	if name == "" {
		return "Hello Bro"
	} else {
		return "Hello" + name
	}
}

func result() {
	result := gethello("Eko")
	fmt.Println(result)
	fmt.Println(gethello(""))
}
