package main

import "fmt"

//1. parameter posisi terakhir bisa dijadikan varags
//2. varags semacam array tapi bisa langsung mengirim datanya, jika lebih dari 1 cukup gunakan koma
//3. jika array wajib membuat array dahulu secara eksplisit sebelum mengirimkan ke function

//variabel argument harus di paling belakang dan tdk bisa dibuat 2
func sumall(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func Variadic_Function() {
	total := sumall(10, 20, 90, 10, 10, 10)
	fmt.Println(total)

	//slice parameter

	numbers := []int{10, 20, 90, 50, 10, 10}
	total = sumall(numbers...)
	fmt.Println(total)
}
