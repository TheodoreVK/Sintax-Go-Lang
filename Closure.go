package main

import "fmt"

// Closure = sebuah function/block spot
// yang bisa interaksi dengan function lain
func Closure() {
	name := "Eko"
	counter := 0
	increment := func() {
		name := "Budi"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}
	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}
