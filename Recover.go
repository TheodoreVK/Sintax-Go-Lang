package main

import "fmt"

//Recover pastikan disimpan di DEFER function
func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message : ", message)
	}
	fmt.Println("Aplikasi Selesai")
}
func runApp(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}

	fmt.Println("Aplikasi Berjalan")
}

func Recover() {
	runApp(false)
	fmt.Println("Eko")
}
