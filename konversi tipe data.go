package main

import "fmt"

func konversi_tipe_data() {
	var nilai32 int32 = 1270000
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name = "Eko Kurniawan"
	var e byte = name[0]
	var estring string = string(e)

	fmt.Println(name)
	fmt.Println(estring)
}