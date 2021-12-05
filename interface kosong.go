package main

import "fmt"

/**
Contoh penggunaan interface kosong
1. fmt.Println(a ...interface{})
2. panic(v interface{})
3. Recover() interface {}
*/

func Ups() interface{} {
	//return 1
	//return true
	return "Ups"
}

func interface_kosong() {
	kosong := Ups()
	fmt.Println(kosong)
}
