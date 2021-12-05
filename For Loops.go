package main

import "fmt"

func perulangan() {
	//Bentuk ke-1
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}
	//Bentuk ke-2
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}
	//bentuk ke-3
	for i := 0; i < 10; i++ {
		fmt.Println("Hi")
	}

	slice := []string{"Eko", "Kurniawan", "Khanedddy", "Budi", "Joko"}

	for i := 0; i < len(slice); i++ {

		fmt.Println(slice[i])
	}
	// For range
	for i, value := range slice {
		fmt.Println("index", i, "=", value)
	}
	book := make(map[string]string)
	book["title"] = "buku Go-Lang"
	book["author"] = "Eko Kurniawan"

	for key, value := range book {
		fmt.Println(key, "=", value)
	}
}
