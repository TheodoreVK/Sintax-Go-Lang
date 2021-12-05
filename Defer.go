package main

import "fmt"

// DEFER = menjalankan function di tahap terakhir meskipun error
func Logging() {
	fmt.Println("Selesai memanggil Function")
}

func runAppLoocation(value int) {
	defer Logging()
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("Result", result)
}

// DEFER = Perintah untuk menjalankan function terakhir
func Defer() {
	runAppLoocation(10)
}
