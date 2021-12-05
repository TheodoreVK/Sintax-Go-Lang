package main

import "fmt"

func getfullname() (string, string, string) {
	return "Eko", "Khaneddy", "Budi"
}

// underscore (_) untuk mengabaikan / ignore return value

func Returning_multiple_values() {
	firstname, middlename, _ := getfullname()
	fmt.Println(firstname, middlename)

}
